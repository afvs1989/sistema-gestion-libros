package outbound

import "sistema-gestion-libros/internal/domain/entity"

// Cuentas es el puerto de persistencia de credenciales de acceso.
type Cuentas interface {
	GuardarCuenta(cuenta *entity.Cuenta) error
	ObtenerPorUsername(username string) (*entity.Cuenta, error)
}

// Hasher es el puerto secundario para el cifrado/verificación de contraseñas.
// La implementación concreta (bcrypt) vive en la capa de infraestructura,
// manteniendo el dominio independiente de la criptografía (DIP).
type Hasher interface {
	Hash(passwordPlano string) (string, error)
	Verificar(hash, passwordPlano string) bool
}
