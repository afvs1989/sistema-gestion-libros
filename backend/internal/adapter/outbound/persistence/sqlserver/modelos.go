// Package sqlserver implementa los puertos de persistencia (outbound) sobre
// SQL Server usando GORM con enfoque "Code First": los modelos definidos en
// código generan el esquema de la base de datos mediante AutoMigrate.
package sqlserver

import "time"

// LibroModel mapea la entidad Libro a la tabla "libros".
type LibroModel struct {
	ID            string `gorm:"primaryKey;size:50"`
	ISBN          string `gorm:"size:30;uniqueIndex;not null"`
	Titulo        string `gorm:"size:200;not null"`
	AutorNombre   string `gorm:"size:100"`
	AutorApellido string `gorm:"size:100"`
	AutorPais     string `gorm:"size:100"`
	Anio          int
	Genero        string `gorm:"size:100"`
	Disponible    bool
	FechaRegistro time.Time
}

func (LibroModel) TableName() string { return "libros" }

// RevistaModel mapea la entidad Revista a la tabla "revistas".
type RevistaModel struct {
	ID         string `gorm:"primaryKey;size:50"`
	Titulo     string `gorm:"size:200;not null"`
	Numero     int
	Editorial  string `gorm:"size:100"`
	Disponible bool
}

func (RevistaModel) TableName() string { return "revistas" }

// UsuarioModel mapea la entidad Usuario (lector) a la tabla "usuarios".
type UsuarioModel struct {
	ID     string `gorm:"primaryKey;size:50"`
	Nombre string `gorm:"size:100;not null"`
	Email  string `gorm:"size:150;uniqueIndex;not null"`
	Activo bool
}

func (UsuarioModel) TableName() string { return "usuarios" }

// PrestamoModel mapea la entidad Prestamo a la tabla "prestamos".
type PrestamoModel struct {
	ID              string `gorm:"primaryKey;size:50"`
	UsuarioID       string `gorm:"size:50;index"`
	RecursoID       string `gorm:"size:50;index"`
	FechaPrestamo   time.Time
	FechaDevolucion *time.Time
	Activo          bool
}

func (PrestamoModel) TableName() string { return "prestamos" }

// CuentaModel mapea la entidad Cuenta (credencial JWT) a la tabla "cuentas".
type CuentaModel struct {
	ID            string `gorm:"primaryKey;size:50"`
	Username      string `gorm:"size:100;uniqueIndex;not null"`
	PasswordHash  string `gorm:"size:255;not null"`
	Rol           string `gorm:"size:50"`
	FechaCreacion time.Time
}

func (CuentaModel) TableName() string { return "cuentas" }
