package errores

import "errors"

var (
	ErrCredencialesInvalidas = errors.New("usuario o contraseña incorrectos")
	ErrCuentaYaExiste        = errors.New("ya existe una cuenta con ese usuario")
	ErrCuentaNoEncontrada    = errors.New("cuenta no encontrada")
)
