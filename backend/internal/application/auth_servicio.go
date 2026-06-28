package application

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/inbound"
	"sistema-gestion-libros/internal/port/outbound"
)

// AuthServicio implementa los casos de uso de autenticación (SRP).
// Depende de puertos secundarios: repositorio de cuentas y hasher (DIP).
type AuthServicio struct {
	cuentas outbound.Cuentas
	hasher  outbound.Hasher
}

var _ inbound.AuthCasosUso = (*AuthServicio)(nil)

func NuevoAuthServicio(cuentas outbound.Cuentas, hasher outbound.Hasher) *AuthServicio {
	return &AuthServicio{cuentas: cuentas, hasher: hasher}
}

// Registrar crea una nueva cuenta cifrando la contraseña antes de persistirla.
func (s *AuthServicio) Registrar(username, password, rol string) (*entity.Cuenta, error) {
	if len(password) < 4 {
		return nil, fmt.Errorf("%w: la contraseña debe tener al menos 4 caracteres", errores.ErrCampoInvalido)
	}
	if existente, _ := s.cuentas.ObtenerPorUsername(username); existente != nil {
		return nil, errores.ErrCuentaYaExiste
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("registrar cuenta: cifrar contraseña: %w", err)
	}

	cuenta, err := entity.NuevaCuenta(GenerarID(), username, hash, rol)
	if err != nil {
		return nil, fmt.Errorf("registrar cuenta: %w", err)
	}
	if err := s.cuentas.GuardarCuenta(cuenta); err != nil {
		return nil, fmt.Errorf("registrar cuenta: %w", err)
	}
	return cuenta, nil
}

// Autenticar verifica las credenciales y devuelve la cuenta si son válidas.
func (s *AuthServicio) Autenticar(username, password string) (*entity.Cuenta, error) {
	cuenta, err := s.cuentas.ObtenerPorUsername(username)
	if err != nil || cuenta == nil {
		return nil, errores.ErrCredencialesInvalidas
	}
	if !s.hasher.Verificar(cuenta.PasswordHash(), password) {
		return nil, errores.ErrCredencialesInvalidas
	}
	return cuenta, nil
}

// GenerarID produce un identificador aleatorio para entidades creadas vía API.
func GenerarID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
