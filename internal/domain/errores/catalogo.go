package errores

import "errors"

var (
	ErrLibroNoEncontrado = errors.New("libro no encontrado")
	ErrLibroNoDisponible = errors.New("el libro no está disponible para préstamo")
	ErrISBNDuplicado     = errors.New("ya existe un libro con ese ISBN")
)
