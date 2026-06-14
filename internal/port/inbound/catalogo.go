package inbound

import "sistema-gestion-libros/internal/domain/entity"

// CatalogoCasosUso agrupa operaciones sobre el catálogo (SRP + ISP).
type CatalogoCasosUso interface {
	RegistrarLibro(libro *entity.Libro) error
	RegistrarRecurso(recurso entity.Recurso) error
	BuscarLibrosPorTitulo(termino string) []*entity.Libro
	BuscarLibrosPorAutor(nombreAutor string) []*entity.Libro
	ListarLibros() []*entity.Libro
	ListarRecursosDisponibles() []entity.Recurso
	EliminarLibro(id string) error
}
