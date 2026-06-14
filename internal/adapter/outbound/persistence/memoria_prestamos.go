package persistence

import (
	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
)

func (m *Memoria) GuardarPrestamo(prestamo *entity.Prestamo) error {
	m.prestamos[prestamo.ID()] = prestamo
	if prestamo.Activo() {
		m.prestamosActivos[prestamo.RecursoID()] = prestamo.ID()
	} else {
		delete(m.prestamosActivos, prestamo.RecursoID())
	}
	return nil
}

func (m *Memoria) ObtenerActivoPorRecurso(recursoID string) (*entity.Prestamo, error) {
	prestamoID, ok := m.prestamosActivos[recursoID]
	if !ok {
		return nil, errores.ErrPrestamoNoEncontrado
	}
	return m.prestamos[prestamoID], nil
}

func (m *Memoria) ListarActivos() []*entity.Prestamo {
	resultado := make([]*entity.Prestamo, 0)
	for _, id := range m.prestamosActivos {
		resultado = append(resultado, m.prestamos[id])
	}
	return resultado
}
