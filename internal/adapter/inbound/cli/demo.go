package cli

import (
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/port/inbound"
)

// CargarDatosDemo inserta un conjunto de registros de ejemplo en el sistema.
func CargarDatosDemo(svc inbound.Biblioteca) error {
	autor1, err := entity.NuevoAutor("Gabriel", "García Márquez", "Colombia")
	if err != nil {
		return err
	}
	autor2, err := entity.NuevoAutor("Isabel", "Allende", "Chile")
	if err != nil {
		return err
	}

	libro1, err := entity.NuevoLibro("L001", "978-0307474728", "Cien años de soledad", autor1, 1967, "Realismo mágico")
	if err != nil {
		return err
	}
	libro2, err := entity.NuevoLibro("L002", "978-0060099455", "La casa de los espíritus", autor2, 1982, "Ficción")
	if err != nil {
		return err
	}

	for _, libro := range []*entity.Libro{libro1, libro2} {
		if err := svc.RegistrarLibro(libro); err != nil {
			return fmt.Errorf("registrar %s: %w", libro.Titulo(), err)
		}
	}

	revista, err := entity.NuevaRevista("R001", "National Geographic", "NatGeo", 245)
	if err != nil {
		return err
	}
	if err := svc.RegistrarRecurso(revista); err != nil {
		return err
	}

	usuario, err := entity.NuevoUsuario("U001", "Ana Pérez", "ana.perez@uide.edu.ec")
	if err != nil {
		return err
	}
	if err := svc.RegistrarUsuario(usuario); err != nil {
		return err
	}

	return nil
}
