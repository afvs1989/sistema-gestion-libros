package outbound

import "sistema-gestion-libros/internal/domain/entity"

// Usuarios es el puerto de persistencia de lectores.
type Usuarios interface {
	GuardarUsuario(usuario *entity.Usuario) error
	ObtenerPorID(id string) (*entity.Usuario, error)
	Listar() []*entity.Usuario
}
