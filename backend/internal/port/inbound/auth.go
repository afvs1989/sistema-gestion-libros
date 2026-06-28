package inbound

import "sistema-gestion-libros/internal/domain/entity"

// AuthCasosUso agrupa las operaciones de autenticación (registro y login).
// El login devuelve la cuenta autenticada; la emisión del token JWT es
// responsabilidad del adaptador inbound HTTP (separación de capas).
type AuthCasosUso interface {
	Registrar(username, password, rol string) (*entity.Cuenta, error)
	Autenticar(username, password string) (*entity.Cuenta, error)
}
