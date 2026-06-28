package sqlserver

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/outbound"
)

// RepoPrestamos implementa el puerto outbound.Prestamos sobre SQL Server.
type RepoPrestamos struct {
	db *gorm.DB
}

var _ outbound.Prestamos = (*RepoPrestamos)(nil)

func NuevoRepoPrestamos(db *gorm.DB) *RepoPrestamos { return &RepoPrestamos{db: db} }

func prestamoAModelo(p *entity.Prestamo) PrestamoModel {
	return PrestamoModel{
		ID:              p.ID(),
		UsuarioID:       p.UsuarioID(),
		RecursoID:       p.RecursoID(),
		FechaPrestamo:   p.FechaPrestamo(),
		FechaDevolucion: p.FechaDevolucion(),
		Activo:          p.Activo(),
	}
}

func modeloAPrestamo(m PrestamoModel) *entity.Prestamo {
	return entity.ReconstruirPrestamo(m.ID, m.UsuarioID, m.RecursoID, m.FechaPrestamo, m.FechaDevolucion, m.Activo)
}

func (r *RepoPrestamos) GuardarPrestamo(prestamo *entity.Prestamo) error {
	if err := r.db.Save(prestamoAModelo(prestamo)).Error; err != nil {
		return fmt.Errorf("guardar préstamo: %w", err)
	}
	return nil
}

func (r *RepoPrestamos) ObtenerActivoPorRecurso(recursoID string) (*entity.Prestamo, error) {
	var m PrestamoModel
	err := r.db.Where("recurso_id = ? AND activo = ?", recursoID, true).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errores.ErrPrestamoNoEncontrado
		}
		return nil, err
	}
	return modeloAPrestamo(m), nil
}

func (r *RepoPrestamos) ListarActivos() []*entity.Prestamo {
	var modelos []PrestamoModel
	r.db.Where("activo = ?", true).Order("fecha_prestamo asc").Find(&modelos)
	prestamos := make([]*entity.Prestamo, 0, len(modelos))
	for _, m := range modelos {
		prestamos = append(prestamos, modeloAPrestamo(m))
	}
	return prestamos
}
