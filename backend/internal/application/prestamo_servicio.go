package application

import (
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/inbound"
	"sistema-gestion-libros/internal/port/outbound"
)

// PrestamoServicio coordina préstamos y devoluciones (SRP).
type PrestamoServicio struct {
	catalogo  outbound.Catalogo
	usuarios  outbound.Usuarios
	prestamos outbound.Prestamos
}

var _ inbound.PrestamoCasosUso = (*PrestamoServicio)(nil)

func NuevoPrestamoServicio(cat outbound.Catalogo, usr outbound.Usuarios, pre outbound.Prestamos) *PrestamoServicio {
	return &PrestamoServicio{catalogo: cat, usuarios: usr, prestamos: pre}
}

func (s *PrestamoServicio) PrestarRecurso(recursoID, usuarioID, prestamoID string) (*entity.Prestamo, error) {
	usuario, err := s.usuarios.ObtenerPorID(usuarioID)
	if err != nil {
		return nil, fmt.Errorf("prestar recurso: %w", err)
	}
	if !usuario.Activo() {
		return nil, fmt.Errorf("prestar recurso: %w: usuario inactivo", errores.ErrCampoInvalido)
	}

	recurso, err := s.catalogo.ObtenerRecursoPorID(recursoID)
	if err != nil {
		return nil, fmt.Errorf("prestar recurso: %w", err)
	}
	if !recurso.Disponible() {
		return nil, fmt.Errorf("prestar recurso: %w", errores.ErrLibroNoDisponible)
	}

	prestamo, err := entity.NuevoPrestamo(prestamoID, usuarioID, recursoID)
	if err != nil {
		return nil, fmt.Errorf("prestar recurso: %w", err)
	}

	recurso.MarcarPrestado()
	if err := s.catalogo.GuardarRecurso(recurso); err != nil {
		return nil, fmt.Errorf("prestar recurso: actualizar recurso: %w", err)
	}
	if err := s.prestamos.GuardarPrestamo(prestamo); err != nil {
		return nil, fmt.Errorf("prestar recurso: guardar préstamo: %w", err)
	}
	return prestamo, nil
}

func (s *PrestamoServicio) DevolverRecurso(recursoID string) error {
	prestamo, err := s.prestamos.ObtenerActivoPorRecurso(recursoID)
	if err != nil {
		return fmt.Errorf("devolver recurso: %w", err)
	}

	recurso, err := s.catalogo.ObtenerRecursoPorID(recursoID)
	if err != nil {
		return fmt.Errorf("devolver recurso: %w", err)
	}

	prestamo.Cerrar()
	recurso.MarcarDevuelto()

	if err := s.prestamos.GuardarPrestamo(prestamo); err != nil {
		return fmt.Errorf("devolver recurso: actualizar préstamo: %w", err)
	}
	if err := s.catalogo.GuardarRecurso(recurso); err != nil {
		return fmt.Errorf("devolver recurso: actualizar recurso: %w", err)
	}
	return nil
}

func (s *PrestamoServicio) ListarPrestamosActivos() []*entity.Prestamo {
	return s.prestamos.ListarActivos()
}
