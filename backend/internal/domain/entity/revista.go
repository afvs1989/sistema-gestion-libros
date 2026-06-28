package entity

import (
	"fmt"
	"strings"

	"sistema-gestion-libros/internal/domain/errores"
)

// Revista es un tipo de recurso catalogable distinto de Libro.
type Revista struct {
	id         string
	titulo     string
	numero     int
	editorial  string
	disponible bool
}

func NuevaRevista(id, titulo, editorial string, numero int) (*Revista, error) {
	if id == "" || strings.TrimSpace(titulo) == "" {
		return nil, fmt.Errorf("%w: id y título son obligatorios", errores.ErrCampoInvalido)
	}
	if numero <= 0 {
		return nil, fmt.Errorf("%w: número de revista inválido", errores.ErrCampoInvalido)
	}
	return &Revista{
		id:         id,
		titulo:     strings.TrimSpace(titulo),
		numero:     numero,
		editorial:  strings.TrimSpace(editorial),
		disponible: true,
	}, nil
}

func (r *Revista) ID() string        { return r.id }
func (r *Revista) Titulo() string    { return r.titulo }
func (r *Revista) Numero() int       { return r.numero }
func (r *Revista) Editorial() string { return r.editorial }
func (r *Revista) Disponible() bool  { return r.disponible }
func (r *Revista) MarcarPrestado()   { r.disponible = false }
func (r *Revista) MarcarDevuelto()   { r.disponible = true }

func (r *Revista) Descripcion() string {
	estado := "disponible"
	if !r.disponible {
		estado = "prestado"
	}
	return fmt.Sprintf("Revista #%d: %s (%s) | %s", r.numero, r.titulo, r.editorial, estado)
}
