package entity

import (
	"fmt"
	"strings"

	"sistema-gestion-libros/internal/domain/errores"
)

// Autor representa al creador de una obra.
type Autor struct {
	nombre   string
	apellido string
	pais     string
}

func NuevoAutor(nombre, apellido, pais string) (*Autor, error) {
	nombre = strings.TrimSpace(nombre)
	apellido = strings.TrimSpace(apellido)
	if nombre == "" || apellido == "" {
		return nil, fmt.Errorf("%w: nombre y apellido son obligatorios", errores.ErrCampoInvalido)
	}
	return &Autor{
		nombre:   nombre,
		apellido: apellido,
		pais:     strings.TrimSpace(pais),
	}, nil
}

func (a *Autor) NombreCompleto() string { return a.nombre + " " + a.apellido }
func (a *Autor) Pais() string           { return a.pais }
