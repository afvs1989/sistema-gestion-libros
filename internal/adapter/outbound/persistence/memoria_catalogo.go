package persistence

import (
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
)

func (m *Memoria) GuardarLibro(libro *entity.Libro) error {
	if _, existe := m.isbnIndex[libro.ISBN()]; existe {
		if existenteID := m.isbnIndex[libro.ISBN()]; existenteID != libro.ID() {
			return errores.ErrISBNDuplicado
		}
	}
	m.libros[libro.ID()] = libro
	m.isbnIndex[libro.ISBN()] = libro.ID()
	m.recursos[libro.ID()] = libro
	return nil
}

func (m *Memoria) ObtenerLibroPorID(id string) (*entity.Libro, error) {
	libro, ok := m.libros[id]
	if !ok {
		return nil, errores.ErrLibroNoEncontrado
	}
	return libro, nil
}

func (m *Memoria) ObtenerLibroPorISBN(isbn string) (*entity.Libro, error) {
	id, ok := m.isbnIndex[isbn]
	if !ok {
		return nil, errores.ErrLibroNoEncontrado
	}
	return m.libros[id], nil
}

func (m *Memoria) ListarLibros() []*entity.Libro {
	resultado := make([]*entity.Libro, 0, len(m.libros))
	for _, libro := range m.libros {
		resultado = append(resultado, libro)
	}
	return resultado
}

func (m *Memoria) EliminarLibro(id string) error {
	libro, ok := m.libros[id]
	if !ok {
		return errores.ErrLibroNoEncontrado
	}
	delete(m.isbnIndex, libro.ISBN())
	delete(m.libros, id)
	delete(m.recursos, id)
	return nil
}

func (m *Memoria) GuardarRecurso(recurso entity.Recurso) error {
	m.recursos[recurso.ID()] = recurso
	return nil
}

func (m *Memoria) ObtenerRecursoPorID(id string) (entity.Recurso, error) {
	recurso, ok := m.recursos[id]
	if !ok {
		return nil, fmt.Errorf("%w: recurso con id %s", errores.ErrLibroNoEncontrado, id)
	}
	return recurso, nil
}

func (m *Memoria) ListarRecursos() []entity.Recurso {
	resultado := make([]entity.Recurso, 0, len(m.recursos))
	for _, r := range m.recursos {
		resultado = append(resultado, r)
	}
	return resultado
}
