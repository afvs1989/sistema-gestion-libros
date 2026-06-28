// Package http es el adaptador inbound que expone los casos de uso como una API
// REST con autenticación JWT y serialización JSON (Gin).
package http

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GestorJWT firma y valida los tokens de acceso.
type GestorJWT struct {
	secreto   []byte
	duracion  time.Duration
	emisor    string
}

// Claims son los datos embebidos en el token.
type Claims struct {
	Username string `json:"username"`
	Rol      string `json:"rol"`
	jwt.RegisteredClaims
}

func NuevoGestorJWT(secreto string, horas int) *GestorJWT {
	return &GestorJWT{
		secreto:  []byte(secreto),
		duracion: time.Duration(horas) * time.Hour,
		emisor:   "sistema-gestion-libros",
	}
}

// Generar firma un token para el usuario y rol indicados.
func (g *GestorJWT) Generar(username, rol string) (string, time.Time, error) {
	expira := time.Now().Add(g.duracion)
	claims := Claims{
		Username: username,
		Rol:      rol,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			Issuer:    g.emisor,
			ExpiresAt: jwt.NewNumericDate(expira),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	firmado, err := token.SignedString(g.secreto)
	return firmado, expira, err
}

// Validar comprueba la firma y vigencia del token y devuelve sus claims.
func (g *GestorJWT) Validar(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inesperado")
		}
		return g.secreto, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido o expirado")
	}
	return claims, nil
}
