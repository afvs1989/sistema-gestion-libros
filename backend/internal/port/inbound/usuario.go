package inbound

import "sistema-gestion-libros/internal/domain/entity"

// UsuarioCasosUso agrupa operaciones sobre lectores.
type UsuarioCasosUso interface {
	RegistrarUsuario(usuario *entity.Usuario) error
	ListarUsuarios() []*entity.Usuario
}
