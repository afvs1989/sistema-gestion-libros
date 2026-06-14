package entity

import (
	"fmt"
	"time"

	"sistema-gestion-libros/internal/domain/errores"
)

// Prestamo registra la relación temporal entre un usuario y un recurso.
type Prestamo struct {
	id              string
	usuarioID       string
	recursoID       string
	fechaPrestamo   time.Time
	fechaDevolucion *time.Time
	activo          bool
}

func NuevoPrestamo(id, usuarioID, recursoID string) (*Prestamo, error) {
	if id == "" || usuarioID == "" || recursoID == "" {
		return nil, fmt.Errorf("%w: id, usuario y recurso son obligatorios", errores.ErrCampoInvalido)
	}
	return &Prestamo{
		id:            id,
		usuarioID:     usuarioID,
		recursoID:     recursoID,
		fechaPrestamo: time.Now(),
		activo:        true,
	}, nil
}

func (p *Prestamo) ID() string               { return p.id }
func (p *Prestamo) UsuarioID() string        { return p.usuarioID }
func (p *Prestamo) RecursoID() string        { return p.recursoID }
func (p *Prestamo) FechaPrestamo() time.Time { return p.fechaPrestamo }
func (p *Prestamo) Activo() bool             { return p.activo }

func (p *Prestamo) Cerrar() {
	ahora := time.Now()
	p.fechaDevolucion = &ahora
	p.activo = false
}

func (p *Prestamo) Resumen() string {
	if p.activo {
		return fmt.Sprintf("Préstamo %s: recurso %s → usuario %s (activo)", p.id, p.recursoID, p.usuarioID)
	}
	return fmt.Sprintf("Préstamo %s: recurso %s → usuario %s (devuelto)", p.id, p.recursoID, p.usuarioID)
}
