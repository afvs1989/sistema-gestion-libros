package application

import (
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/port/inbound"
	"sistema-gestion-libros/internal/port/outbound"
)

// UsuarioServicio gestiona el registro y listado de lectores (SRP).
type UsuarioServicio struct {
	repo outbound.Usuarios
}

var _ inbound.UsuarioCasosUso = (*UsuarioServicio)(nil)

func NuevoUsuarioServicio(repo outbound.Usuarios) *UsuarioServicio {
	return &UsuarioServicio{repo: repo}
}

func (s *UsuarioServicio) RegistrarUsuario(usuario *entity.Usuario) error {
	if err := s.repo.GuardarUsuario(usuario); err != nil {
		return fmt.Errorf("registrar usuario: %w", err)
	}
	return nil
}

func (s *UsuarioServicio) ListarUsuarios() []*entity.Usuario {
	return s.repo.Listar()
}
