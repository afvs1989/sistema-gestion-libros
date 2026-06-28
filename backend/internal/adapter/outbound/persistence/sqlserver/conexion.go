package sqlserver

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config agrupa los parámetros de conexión a SQL Server.
type Config struct {
	Host     string
	Puerto   string
	Usuario  string
	Password string
	Base     string
}

// DSN construye la cadena de conexión de GORM para SQL Server.
func (c Config) DSN() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		c.Usuario, c.Password, c.Host, c.Puerto, c.Base)
}

// Conectar abre la conexión y ejecuta la migración Code First (AutoMigrate),
// creando/actualizando las tablas a partir de los modelos definidos en código.
func Conectar(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("conectar a SQL Server: %w", err)
	}

	if err := db.AutoMigrate(
		&LibroModel{},
		&RevistaModel{},
		&UsuarioModel{},
		&PrestamoModel{},
		&CuentaModel{},
	); err != nil {
		return nil, fmt.Errorf("migrar esquema (Code First): %w", err)
	}
	return db, nil
}
