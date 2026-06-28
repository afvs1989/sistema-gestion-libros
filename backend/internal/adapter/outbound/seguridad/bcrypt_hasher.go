// Package seguridad implementa el puerto outbound.Hasher usando bcrypt.
package seguridad

import (
	"golang.org/x/crypto/bcrypt"

	"sistema-gestion-libros/internal/port/outbound"
)

// BcryptHasher cifra y verifica contraseñas con el algoritmo bcrypt.
type BcryptHasher struct {
	coste int
}

var _ outbound.Hasher = (*BcryptHasher)(nil)

// NuevoBcryptHasher crea el hasher con el coste por defecto recomendado.
func NuevoBcryptHasher() *BcryptHasher {
	return &BcryptHasher{coste: bcrypt.DefaultCost}
}

func (h *BcryptHasher) Hash(passwordPlano string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordPlano), h.coste)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (h *BcryptHasher) Verificar(hash, passwordPlano string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordPlano)) == nil
}
