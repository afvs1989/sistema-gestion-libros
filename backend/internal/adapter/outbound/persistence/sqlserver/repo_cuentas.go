package sqlserver

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/outbound"
)

// RepoCuentas implementa el puerto outbound.Cuentas sobre SQL Server.
type RepoCuentas struct {
	db *gorm.DB
}

var _ outbound.Cuentas = (*RepoCuentas)(nil)

func NuevoRepoCuentas(db *gorm.DB) *RepoCuentas { return &RepoCuentas{db: db} }

func cuentaAModelo(c *entity.Cuenta) CuentaModel {
	return CuentaModel{
		ID:            c.ID(),
		Username:      c.Username(),
		PasswordHash:  c.PasswordHash(),
		Rol:           c.Rol(),
		FechaCreacion: c.FechaCreacion(),
	}
}

func modeloACuenta(m CuentaModel) *entity.Cuenta {
	return entity.ReconstruirCuenta(m.ID, m.Username, m.PasswordHash, m.Rol, m.FechaCreacion)
}

func (r *RepoCuentas) GuardarCuenta(cuenta *entity.Cuenta) error {
	if err := r.db.Save(cuentaAModelo(cuenta)).Error; err != nil {
		return fmt.Errorf("guardar cuenta: %w", err)
	}
	return nil
}

func (r *RepoCuentas) ObtenerPorUsername(username string) (*entity.Cuenta, error) {
	var m CuentaModel
	err := r.db.Where("username = ?", strings.ToLower(strings.TrimSpace(username))).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errores.ErrCuentaNoEncontrada
		}
		return nil, err
	}
	return modeloACuenta(m), nil
}
