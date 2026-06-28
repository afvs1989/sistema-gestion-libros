package entity

import (
	"fmt"
	"strings"
	"time"

	"sistema-gestion-libros/internal/domain/errores"
)

// Roles disponibles para una cuenta del sistema.
const (
	RolAdmin         = "ADMIN"
	RolBibliotecario = "BIBLIOTECARIO"
)

// Cuenta representa una credencial de acceso al sistema (autenticación JWT).
// El dominio nunca conoce la contraseña en claro: solo almacena su hash.
type Cuenta struct {
	id            string
	username      string
	passwordHash  string
	rol           string
	fechaCreacion time.Time
}

// NuevaCuenta crea una credencial nueva. El hash debe calcularse en la capa de
// infraestructura (puerto Hasher) antes de invocar este constructor.
func NuevaCuenta(id, username, passwordHash, rol string) (*Cuenta, error) {
	username = strings.TrimSpace(strings.ToLower(username))
	if id == "" || username == "" || passwordHash == "" {
		return nil, fmt.Errorf("%w: id, usuario y contraseña son obligatorios", errores.ErrCampoInvalido)
	}
	if rol != RolAdmin && rol != RolBibliotecario {
		rol = RolBibliotecario
	}
	return &Cuenta{
		id:            id,
		username:      username,
		passwordHash:  passwordHash,
		rol:           rol,
		fechaCreacion: time.Now(),
	}, nil
}

// ReconstruirCuenta rehidrata una Cuenta desde la persistencia.
func ReconstruirCuenta(id, username, passwordHash, rol string, fechaCreacion time.Time) *Cuenta {
	return &Cuenta{
		id:            id,
		username:      username,
		passwordHash:  passwordHash,
		rol:           rol,
		fechaCreacion: fechaCreacion,
	}
}

func (c *Cuenta) ID() string               { return c.id }
func (c *Cuenta) Username() string         { return c.username }
func (c *Cuenta) PasswordHash() string     { return c.passwordHash }
func (c *Cuenta) Rol() string              { return c.rol }
func (c *Cuenta) FechaCreacion() time.Time { return c.fechaCreacion }
