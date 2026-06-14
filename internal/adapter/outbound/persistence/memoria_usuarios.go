package persistence

import (
	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
)

func (m *Memoria) GuardarUsuario(usuario *entity.Usuario) error {
	m.usuarios[usuario.ID()] = usuario
	return nil
}

func (m *Memoria) ObtenerPorID(id string) (*entity.Usuario, error) {
	usuario, ok := m.usuarios[id]
	if !ok {
		return nil, errores.ErrUsuarioNoEncontrado
	}
	return usuario, nil
}

func (m *Memoria) Listar() []*entity.Usuario {
	resultado := make([]*entity.Usuario, 0, len(m.usuarios))
	for _, u := range m.usuarios {
		resultado = append(resultado, u)
	}
	return resultado
}
