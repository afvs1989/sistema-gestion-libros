package outbound

import "sistema-gestion-libros/internal/domain/entity"

// Catalogo es el puerto de persistencia del catálogo bibliográfico.
type Catalogo interface {
	GuardarLibro(libro *entity.Libro) error
	ObtenerLibroPorID(id string) (*entity.Libro, error)
	ObtenerLibroPorISBN(isbn string) (*entity.Libro, error)
	ListarLibros() []*entity.Libro
	EliminarLibro(id string) error
	GuardarRecurso(recurso entity.Recurso) error
	ObtenerRecursoPorID(id string) (entity.Recurso, error)
	ListarRecursos() []entity.Recurso
}
