package application

import (
	"fmt"

	"sistema-gestion-libros/internal/port/inbound"
	"sistema-gestion-libros/internal/port/outbound"
)

// ConsultaServicio genera informes del estado del sistema (SRP).
type ConsultaServicio struct {
	nombre    string
	catalogo  outbound.Catalogo
	prestamos outbound.Prestamos
	catalogoUC inbound.CatalogoCasosUso
}

var _ inbound.ConsultaBiblioteca = (*ConsultaServicio)(nil)

func NuevoConsultaServicio(
	nombre string,
	cat outbound.Catalogo,
	pre outbound.Prestamos,
	catalogoUC inbound.CatalogoCasosUso,
) *ConsultaServicio {
	return &ConsultaServicio{
		nombre:     nombre,
		catalogo:   cat,
		prestamos:  pre,
		catalogoUC: catalogoUC,
	}
}

func (s *ConsultaServicio) Nombre() string {
	return s.nombre
}

func (s *ConsultaServicio) ResumenCatalogo() string {
	return fmt.Sprintf(
		"Biblioteca '%s' | Libros: %d | Recursos totales: %d | Disponibles: %d | Préstamos activos: %d",
		s.nombre,
		len(s.catalogo.ListarLibros()),
		len(s.catalogo.ListarRecursos()),
		len(s.catalogoUC.ListarRecursosDisponibles()),
		len(s.prestamos.ListarActivos()),
	)
}
