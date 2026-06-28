package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sistema-gestion-libros/internal/application"
	"sistema-gestion-libros/internal/domain/entity"
)

// listarUsuarios godoc
//
//	@Summary		Listar usuarios
//	@Description	Devuelve los lectores registrados.
//	@Tags			Usuarios
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		UsuarioDTO
//	@Failure		401	{object}	map[string]string
//	@Router			/usuarios [get]
func (s *Servidor) listarUsuarios(c *gin.Context) {
	usuarios := s.gestor.ListarUsuarios()
	out := make([]UsuarioDTO, 0, len(usuarios))
	for _, u := range usuarios {
		out = append(out, aUsuarioDTO(u))
	}
	c.JSON(http.StatusOK, out)
}

// crearUsuario godoc
//
//	@Summary		Registrar usuario
//	@Description	Crea un lector nuevo.
//	@Tags			Usuarios
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			usuario	body		CrearUsuarioRequest	true	"Datos del usuario"
//	@Success		201		{object}	UsuarioDTO
//	@Failure		400		{object}	map[string]string
//	@Router			/usuarios [post]
func (s *Servidor) crearUsuario(c *gin.Context) {
	var req CrearUsuarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	usuario, err := entity.NuevoUsuario(application.GenerarID(), req.Nombre, req.Email)
	if err != nil {
		responderError(c, err)
		return
	}
	if err := s.gestor.RegistrarUsuario(usuario); err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusCreated, aUsuarioDTO(usuario))
}
