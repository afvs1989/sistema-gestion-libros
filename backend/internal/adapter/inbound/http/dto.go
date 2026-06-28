package http

import (
	"time"

	"sistema-gestion-libros/internal/domain/entity"
)

// ---- Peticiones (entrada JSON) ----

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegistroRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Rol      string `json:"rol"`
}

type CrearLibroRequest struct {
	ISBN          string `json:"isbn" binding:"required"`
	Titulo        string `json:"titulo" binding:"required"`
	AutorNombre   string `json:"autorNombre" binding:"required"`
	AutorApellido string `json:"autorApellido" binding:"required"`
	AutorPais     string `json:"autorPais"`
	Anio          int    `json:"anio" binding:"required"`
	Genero        string `json:"genero"`
}

type CrearRevistaRequest struct {
	Titulo    string `json:"titulo" binding:"required"`
	Editorial string `json:"editorial"`
	Numero    int    `json:"numero" binding:"required"`
}

type CrearUsuarioRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

type PrestamoRequest struct {
	RecursoID string `json:"recursoId" binding:"required"`
	UsuarioID string `json:"usuarioId" binding:"required"`
}

// ---- Respuestas (salida JSON) ----

type TokenResponse struct {
	Token    string    `json:"token"`
	Tipo     string    `json:"tipo"`
	Username string    `json:"username"`
	Rol      string    `json:"rol"`
	Expira   time.Time `json:"expira"`
}

type AutorDTO struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Pais     string `json:"pais"`
}

type LibroDTO struct {
	ID            string    `json:"id"`
	ISBN          string    `json:"isbn"`
	Titulo        string    `json:"titulo"`
	Autor         AutorDTO  `json:"autor"`
	Anio          int       `json:"anio"`
	Genero        string    `json:"genero"`
	Disponible    bool      `json:"disponible"`
	FechaRegistro time.Time `json:"fechaRegistro"`
}

type RecursoDTO struct {
	ID          string `json:"id"`
	Titulo      string `json:"titulo"`
	Disponible  bool   `json:"disponible"`
	Descripcion string `json:"descripcion"`
}

type UsuarioDTO struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Activo bool   `json:"activo"`
}

type PrestamoDTO struct {
	ID              string     `json:"id"`
	UsuarioID       string     `json:"usuarioId"`
	RecursoID       string     `json:"recursoId"`
	FechaPrestamo   time.Time  `json:"fechaPrestamo"`
	FechaDevolucion *time.Time `json:"fechaDevolucion"`
	Activo          bool       `json:"activo"`
}

// ---- Mapeo entidad -> DTO ----

func aLibroDTO(l *entity.Libro) LibroDTO {
	return LibroDTO{
		ID:     l.ID(),
		ISBN:   l.ISBN(),
		Titulo: l.Titulo(),
		Autor: AutorDTO{
			Nombre:   l.Autor().Nombre(),
			Apellido: l.Autor().Apellido(),
			Pais:     l.Autor().Pais(),
		},
		Anio:          l.Anio(),
		Genero:        l.Genero(),
		Disponible:    l.Disponible(),
		FechaRegistro: l.FechaRegistro(),
	}
}

func aLibrosDTO(libros []*entity.Libro) []LibroDTO {
	out := make([]LibroDTO, 0, len(libros))
	for _, l := range libros {
		out = append(out, aLibroDTO(l))
	}
	return out
}

func aRecursoDTO(r entity.Recurso) RecursoDTO {
	return RecursoDTO{
		ID:          r.ID(),
		Titulo:      r.Titulo(),
		Disponible:  r.Disponible(),
		Descripcion: r.Descripcion(),
	}
}

func aUsuarioDTO(u *entity.Usuario) UsuarioDTO {
	return UsuarioDTO{ID: u.ID(), Nombre: u.Nombre(), Email: u.Email(), Activo: u.Activo()}
}

func aPrestamoDTO(p *entity.Prestamo) PrestamoDTO {
	return PrestamoDTO{
		ID:              p.ID(),
		UsuarioID:       p.UsuarioID(),
		RecursoID:       p.RecursoID(),
		FechaPrestamo:   p.FechaPrestamo(),
		FechaDevolucion: p.FechaDevolucion(),
		Activo:          p.Activo(),
	}
}
