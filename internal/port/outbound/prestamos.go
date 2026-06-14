package outbound

import "sistema-gestion-libros/internal/domain/entity"

// Prestamos es el puerto de persistencia de préstamos.
type Prestamos interface {
	GuardarPrestamo(prestamo *entity.Prestamo) error
	ObtenerActivoPorRecurso(recursoID string) (*entity.Prestamo, error)
	ListarActivos() []*entity.Prestamo
}
