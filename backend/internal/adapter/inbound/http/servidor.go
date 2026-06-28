package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/inbound"
)

// Servidor agrupa las dependencias del adaptador HTTP (casos de uso + JWT).
type Servidor struct {
	gestor inbound.Biblioteca
	auth   inbound.AuthCasosUso
	jwt    *GestorJWT
}

func NuevoServidor(gestor inbound.Biblioteca, auth inbound.AuthCasosUso, jwt *GestorJWT) *Servidor {
	return &Servidor{gestor: gestor, auth: auth, jwt: jwt}
}

// Router construye el motor Gin con CORS, rutas públicas y rutas protegidas por JWT.
func (s *Servidor) Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), cors())

	// Documentación interactiva (Swagger UI) en /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	// Público.
	api.GET("/health", s.health)
	api.POST("/auth/register", s.registrar)
	api.POST("/auth/login", s.login)

	// Protegido con JWT.
	priv := api.Group("/")
	priv.Use(s.requiereJWT())
	{
		priv.GET("/libros", s.listarLibros)
		priv.POST("/libros", s.crearLibro)
		priv.GET("/libros/buscar", s.buscarLibros)
		priv.GET("/libros/:id", s.obtenerLibro)
		priv.DELETE("/libros/:id", s.eliminarLibro)

		priv.POST("/revistas", s.crearRevista)
		priv.GET("/recursos/disponibles", s.recursosDisponibles)

		priv.GET("/usuarios", s.listarUsuarios)
		priv.POST("/usuarios", s.crearUsuario)

		priv.POST("/prestamos", s.prestar)
		priv.PUT("/prestamos/devolver/:recursoId", s.devolver)
		priv.GET("/prestamos/activos", s.prestamosActivos)

		priv.GET("/catalogo/resumen", s.resumen)
	}

	return r
}

// requiereJWT es el middleware que valida el header Authorization: Bearer <token>.
func (s *Servidor) requiereJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
			return
		}
		claims, err := s.jwt.Validar(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("username", claims.Username)
		c.Set("rol", claims.Rol)
		c.Next()
	}
}

// cors habilita peticiones desde el frontend Angular (otro origen).
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// health godoc
//
//	@Summary	Estado del servicio
//	@Tags		Sistema
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Router		/health [get]
func (s *Servidor) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"estado": "ok", "servicio": "sistema-gestion-libros"})
}

// responderError traduce errores de dominio a códigos HTTP adecuados.
func responderError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errores.ErrCredencialesInvalidas):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case errors.Is(err, errores.ErrCuentaYaExiste), errors.Is(err, errores.ErrISBNDuplicado):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, errores.ErrLibroNoEncontrado),
		errors.Is(err, errores.ErrUsuarioNoEncontrado),
		errors.Is(err, errores.ErrPrestamoNoEncontrado),
		errors.Is(err, errores.ErrCuentaNoEncontrada):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, errores.ErrCampoInvalido), errors.Is(err, errores.ErrLibroNoDisponible):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
