package sqlserver

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/outbound"
)

// RepoUsuarios implementa el puerto outbound.Usuarios sobre SQL Server.
type RepoUsuarios struct {
	db *gorm.DB
}

var _ outbound.Usuarios = (*RepoUsuarios)(nil)

func NuevoRepoUsuarios(db *gorm.DB) *RepoUsuarios { return &RepoUsuarios{db: db} }

func usuarioAModelo(u *entity.Usuario) UsuarioModel {
	return UsuarioModel{ID: u.ID(), Nombre: u.Nombre(), Email: u.Email(), Activo: u.Activo()}
}

func modeloAUsuario(m UsuarioModel) *entity.Usuario {
	return entity.ReconstruirUsuario(m.ID, m.Nombre, m.Email, m.Activo)
}

func (r *RepoUsuarios) GuardarUsuario(usuario *entity.Usuario) error {
	if err := r.db.Save(usuarioAModelo(usuario)).Error; err != nil {
		return fmt.Errorf("guardar usuario: %w", err)
	}
	return nil
}

func (r *RepoUsuarios) ObtenerPorID(id string) (*entity.Usuario, error) {
	var m UsuarioModel
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errores.ErrUsuarioNoEncontrado
		}
		return nil, err
	}
	return modeloAUsuario(m), nil
}

func (r *RepoUsuarios) Listar() []*entity.Usuario {
	var modelos []UsuarioModel
	r.db.Order("nombre asc").Find(&modelos)
	usuarios := make([]*entity.Usuario, 0, len(modelos))
	for _, m := range modelos {
		usuarios = append(usuarios, modeloAUsuario(m))
	}
	return usuarios
}
