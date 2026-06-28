package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sistema-gestion-libros/internal/application"
)

// prestar godoc
//
//	@Summary		Prestar recurso
//	@Description	Registra el préstamo de un recurso (libro/revista) a un usuario.
//	@Tags			Préstamos
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			prestamo	body		PrestamoRequest	true	"IDs de recurso y usuario"
//	@Success		201			{object}	PrestamoDTO
//	@Failure		400			{object}	map[string]string
//	@Failure		404			{object}	map[string]string
//	@Router			/prestamos [post]
func (s *Servidor) prestar(c *gin.Context) {
	var req PrestamoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	prestamo, err := s.gestor.PrestarRecurso(req.RecursoID, req.UsuarioID, application.GenerarID())
	if err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusCreated, aPrestamoDTO(prestamo))
}

// devolver godoc
//
//	@Summary		Devolver recurso
//	@Description	Cierra el préstamo activo asociado a un recurso.
//	@Tags			Préstamos
//	@Produce		json
//	@Security		BearerAuth
//	@Param			recursoId	path		string	true	"ID del recurso a devolver"
//	@Success		200			{object}	map[string]string
//	@Failure		404			{object}	map[string]string
//	@Router			/prestamos/devolver/{recursoId} [put]
func (s *Servidor) devolver(c *gin.Context) {
	if err := s.gestor.DevolverRecurso(c.Param("recursoId")); err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "recurso devuelto"})
}

// prestamosActivos godoc
//
//	@Summary		Préstamos activos
//	@Description	Lista los préstamos que siguen abiertos (sin devolver).
//	@Tags			Préstamos
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		PrestamoDTO
//	@Failure		401	{object}	map[string]string
//	@Router			/prestamos/activos [get]
func (s *Servidor) prestamosActivos(c *gin.Context) {
	prestamos := s.gestor.ListarPrestamosActivos()
	out := make([]PrestamoDTO, 0, len(prestamos))
	for _, p := range prestamos {
		out = append(out, aPrestamoDTO(p))
	}
	c.JSON(http.StatusOK, out)
}

// resumen godoc
//
//	@Summary		Resumen del catálogo
//	@Description	Informe agregado: totales de libros, recursos, disponibles y préstamos activos.
//	@Tags			Catálogo
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	map[string]string
//	@Router			/catalogo/resumen [get]
func (s *Servidor) resumen(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"biblioteca": s.gestor.Nombre(),
		"resumen":    s.gestor.ResumenCatalogo(),
	})
}
