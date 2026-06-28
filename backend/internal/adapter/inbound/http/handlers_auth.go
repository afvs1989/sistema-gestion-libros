package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// registrar godoc
//
//	@Summary		Registrar cuenta de acceso
//	@Description	Crea una nueva cuenta (credencial) cifrando la contraseña con bcrypt.
//	@Tags			Autenticación
//	@Accept			json
//	@Produce		json
//	@Param			cuenta	body		RegistroRequest	true	"Datos de la cuenta"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		409		{object}	map[string]string
//	@Router			/auth/register [post]
func (s *Servidor) registrar(c *gin.Context) {
	var req RegistroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	cuenta, err := s.auth.Registrar(req.Username, req.Password, req.Rol)
	if err != nil {
		responderError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       cuenta.ID(),
		"username": cuenta.Username(),
		"rol":      cuenta.Rol(),
	})
}

// login godoc
//
//	@Summary		Iniciar sesión (emite token JWT)
//	@Description	Verifica las credenciales y devuelve un token JWT para autenticar el resto de servicios.
//	@Tags			Autenticación
//	@Accept			json
//	@Produce		json
//	@Param			credenciales	body		LoginRequest	true	"Usuario y contraseña"
//	@Success		200				{object}	TokenResponse
//	@Failure		400				{object}	map[string]string
//	@Failure		401				{object}	map[string]string
//	@Router			/auth/login [post]
func (s *Servidor) login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos: " + err.Error()})
		return
	}
	cuenta, err := s.auth.Autenticar(req.Username, req.Password)
	if err != nil {
		responderError(c, err)
		return
	}
	token, expira, err := s.jwt.Generar(cuenta.Username(), cuenta.Rol())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo emitir el token"})
		return
	}
	c.JSON(http.StatusOK, TokenResponse{
		Token:    token,
		Tipo:     "Bearer",
		Username: cuenta.Username(),
		Rol:      cuenta.Rol(),
		Expira:   expira,
	})
}
