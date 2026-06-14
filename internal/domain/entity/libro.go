package entity

import (
	"fmt"
	"strings"
	"time"

	"sistema-gestion-libros/internal/domain/errores"
)

// Libro representa una obra impresa catalogada en la biblioteca.
type Libro struct {
	id            string
	isbn          string
	titulo        string
	autor         *Autor
	anio          int
	genero        string
	disponible    bool
	fechaRegistro time.Time
}

func NuevoLibro(id, isbn, titulo string, autor *Autor, anio int, genero string) (*Libro, error) {
	isbn = strings.TrimSpace(isbn)
	titulo = strings.TrimSpace(titulo)
	if id == "" || isbn == "" || titulo == "" {
		return nil, fmt.Errorf("%w: id, isbn y título son obligatorios", errores.ErrCampoInvalido)
	}
	if autor == nil {
		return nil, fmt.Errorf("%w: el libro debe tener un autor", errores.ErrCampoInvalido)
	}
	if anio < 1000 || anio > time.Now().Year() {
		return nil, fmt.Errorf("%w: año de publicación inválido", errores.ErrCampoInvalido)
	}
	return &Libro{
		id:            id,
		isbn:          isbn,
		titulo:        titulo,
		autor:         autor,
		anio:          anio,
		genero:        strings.TrimSpace(genero),
		disponible:    true,
		fechaRegistro: time.Now(),
	}, nil
}

func (l *Libro) ID() string               { return l.id }
func (l *Libro) ISBN() string             { return l.isbn }
func (l *Libro) Titulo() string           { return l.titulo }
func (l *Libro) Autor() *Autor            { return l.autor }
func (l *Libro) Anio() int                { return l.anio }
func (l *Libro) Genero() string           { return l.genero }
func (l *Libro) Disponible() bool         { return l.disponible }
func (l *Libro) FechaRegistro() time.Time { return l.fechaRegistro }
func (l *Libro) MarcarPrestado()          { l.disponible = false }
func (l *Libro) MarcarDevuelto()          { l.disponible = true }

func (l *Libro) Descripcion() string {
	estado := "disponible"
	if !l.disponible {
		estado = "prestado"
	}
	return fmt.Sprintf("[%s] %s - %s (%d) | %s", l.isbn, l.titulo, l.autor.NombreCompleto(), l.anio, estado)
}
