package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sistema-gestion-libros/internal/application"
	"sistema-gestion-libros/internal/domain/entity"
)

// listarLibros godoc
//
//	@Summary		Listar libros
//	@Description	Devuelve el catálogo completo de libros.
//	@Tags			Libros
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		LibroDTO
//	@Failure		401	{object}	map[string]string
//	@Router			/libros [get]
func (s *Servidor) listarLibros(c *gin.Context) {
	c.JSON(http.StatusOK, aLibrosDTO(s.gestor.ListarLibros()))
}

// crearLibro godoc
//
//	@Summary		Registrar libro
//	@Description	Crea un libro nuevo en el catálogo.
//	@Tags			Libros
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			libro	body		CrearLibroRequest	true	"Datos del libro"
//	@Success		201		{object}	LibroDTO
//	@Failure		400		{object}	map[string]string
//	@Failure		409		{object}	map[string]string
//	@Router			/libros [post]
func (s *Servidor) crearLibro(c *gin.Context) {
	var req CrearLibroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	autor, err := entity.NuevoAutor(req.AutorNombre, req.AutorApellido, req.AutorPais)
	if err != nil {
		responderError(c, err)
		return
	}
	libro, err := entity.NuevoLibro(application.GenerarID(), req.ISBN, req.Titulo, autor, req.Anio, req.Genero)
	if err != nil {
		responderError(c, err)
		return
	}
	if err := s.gestor.RegistrarLibro(libro); err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusCreated, aLibroDTO(libro))
}

// buscarLibros godoc
//
//	@Summary		Buscar libros
//	@Description	Filtra libros por título o por autor (usa "autor" si viene; si no, "titulo").
//	@Tags			Libros
//	@Produce		json
//	@Security		BearerAuth
//	@Param			titulo	query		string	false	"Texto a buscar en el título"
//	@Param			autor	query		string	false	"Texto a buscar en el autor"
//	@Success		200		{array}		LibroDTO
//	@Failure		401		{object}	map[string]string
//	@Router			/libros/buscar [get]
func (s *Servidor) buscarLibros(c *gin.Context) {
	if autor := c.Query("autor"); autor != "" {
		c.JSON(http.StatusOK, aLibrosDTO(s.gestor.BuscarLibrosPorAutor(autor)))
		return
	}
	titulo := c.Query("titulo")
	c.JSON(http.StatusOK, aLibrosDTO(s.gestor.BuscarLibrosPorTitulo(titulo)))
}

// obtenerLibro godoc
//
//	@Summary		Obtener libro por ID
//	@Tags			Libros
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID del libro"
//	@Success		200	{object}	LibroDTO
//	@Failure		404	{object}	map[string]string
//	@Router			/libros/{id} [get]
func (s *Servidor) obtenerLibro(c *gin.Context) {
	for _, l := range s.gestor.ListarLibros() {
		if l.ID() == c.Param("id") {
			c.JSON(http.StatusOK, aLibroDTO(l))
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "libro no encontrado"})
}

// eliminarLibro godoc
//
//	@Summary		Eliminar libro
//	@Tags			Libros
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID del libro"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/libros/{id} [delete]
func (s *Servidor) eliminarLibro(c *gin.Context) {
	if err := s.gestor.EliminarLibro(c.Param("id")); err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "libro eliminado"})
}

// crearRevista godoc
//
//	@Summary		Registrar revista
//	@Description	Crea una revista (otro tipo de Recurso catalogable).
//	@Tags			Catálogo
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			revista	body		CrearRevistaRequest	true	"Datos de la revista"
//	@Success		201		{object}	RecursoDTO
//	@Failure		400		{object}	map[string]string
//	@Router			/revistas [post]
func (s *Servidor) crearRevista(c *gin.Context) {
	var req CrearRevistaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	revista, err := entity.NuevaRevista(application.GenerarID(), req.Titulo, req.Editorial, req.Numero)
	if err != nil {
		responderError(c, err)
		return
	}
	if err := s.gestor.RegistrarRecurso(revista); err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusCreated, aRecursoDTO(revista))
}

// recursosDisponibles godoc
//
//	@Summary		Recursos disponibles
//	@Description	Lista libros y revistas disponibles para préstamo (polimorfismo Recurso).
//	@Tags			Catálogo
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{array}		RecursoDTO
//	@Failure		401	{object}	map[string]string
//	@Router			/recursos/disponibles [get]
func (s *Servidor) recursosDisponibles(c *gin.Context) {
	recursos := s.gestor.ListarRecursosDisponibles()
	out := make([]RecursoDTO, 0, len(recursos))
	for _, r := range recursos {
		out = append(out, aRecursoDTO(r))
	}
	c.JSON(http.StatusOK, out)
}
