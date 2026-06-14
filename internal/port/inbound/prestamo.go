package inbound

import "sistema-gestion-libros/internal/domain/entity"

// PrestamoCasosUso agrupa operaciones de préstamo y devolución.
type PrestamoCasosUso interface {
	PrestarRecurso(recursoID, usuarioID, prestamoID string) (*entity.Prestamo, error)
	DevolverRecurso(recursoID string) error
	ListarPrestamosActivos() []*entity.Prestamo
}
