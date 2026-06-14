package application

import (
	"fmt"
	"strings"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/port/inbound"
	"sistema-gestion-libros/internal/port/outbound"
)

// CatalogoServicio gestiona registro y consulta del catálogo (SRP).
type CatalogoServicio struct {
	repo outbound.Catalogo
}

var _ inbound.CatalogoCasosUso = (*CatalogoServicio)(nil)

func NuevoCatalogoServicio(repo outbound.Catalogo) *CatalogoServicio {
	return &CatalogoServicio{repo: repo}
}

func (s *CatalogoServicio) RegistrarLibro(libro *entity.Libro) error {
	if err := s.repo.GuardarLibro(libro); err != nil {
		return fmt.Errorf("registrar libro: %w", err)
	}
	return nil
}

func (s *CatalogoServicio) RegistrarRecurso(recurso entity.Recurso) error {
	if err := s.repo.GuardarRecurso(recurso); err != nil {
		return fmt.Errorf("registrar recurso: %w", err)
	}
	return nil
}

func (s *CatalogoServicio) BuscarLibrosPorTitulo(termino string) []*entity.Libro {
	termino = strings.ToLower(strings.TrimSpace(termino))
	if termino == "" {
		return s.repo.ListarLibros()
	}
	var coincidencias []*entity.Libro
	for _, libro := range s.repo.ListarLibros() {
		if strings.Contains(strings.ToLower(libro.Titulo()), termino) {
			coincidencias = append(coincidencias, libro)
		}
	}
	return coincidencias
}

func (s *CatalogoServicio) BuscarLibrosPorAutor(nombreAutor string) []*entity.Libro {
	nombreAutor = strings.ToLower(strings.TrimSpace(nombreAutor))
	var coincidencias []*entity.Libro
	for _, libro := range s.repo.ListarLibros() {
		if strings.Contains(strings.ToLower(libro.Autor().NombreCompleto()), nombreAutor) {
			coincidencias = append(coincidencias, libro)
		}
	}
	return coincidencias
}

func (s *CatalogoServicio) ListarLibros() []*entity.Libro {
	return s.repo.ListarLibros()
}

func (s *CatalogoServicio) ListarRecursosDisponibles() []entity.Recurso {
	var disponibles []entity.Recurso
	for _, recurso := range s.repo.ListarRecursos() {
		if recurso.Disponible() {
			disponibles = append(disponibles, recurso)
		}
	}
	return disponibles
}

func (s *CatalogoServicio) EliminarLibro(id string) error {
	if err := s.repo.EliminarLibro(id); err != nil {
		return fmt.Errorf("eliminar libro: %w", err)
	}
	return nil
}
