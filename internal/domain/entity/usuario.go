package entity

import (
	"fmt"
	"strings"

	"sistema-gestion-libros/internal/domain/errores"
)

// Usuario representa a un lector registrado en la biblioteca.
type Usuario struct {
	id     string
	nombre string
	email  string
	activo bool
}

func NuevoUsuario(id, nombre, email string) (*Usuario, error) {
	nombre = strings.TrimSpace(nombre)
	email = strings.TrimSpace(email)
	if id == "" || nombre == "" || email == "" {
		return nil, fmt.Errorf("%w: id, nombre y email son obligatorios", errores.ErrCampoInvalido)
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("%w: formato de email inválido", errores.ErrCampoInvalido)
	}
	return &Usuario{id: id, nombre: nombre, email: email, activo: true}, nil
}

func (u *Usuario) ID() string     { return u.id }
func (u *Usuario) Nombre() string { return u.nombre }
func (u *Usuario) Email() string  { return u.email }
func (u *Usuario) Activo() bool   { return u.activo }
