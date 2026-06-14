// Punto de entrada: ensambla adaptadores y casos de uso (arquitectura hexagonal + SOLID).
package main

import (
	"fmt"

	"sistema-gestion-libros/internal/adapter/inbound/cli"
	"sistema-gestion-libros/internal/adapter/outbound/persistence"
	"sistema-gestion-libros/internal/application"
)

func main() {
	fmt.Println("=== Sistema de Gestión de Libros - Avance Autónomo 2 ===")

	repo := persistence.NuevaMemoria()

	catalogoSvc := application.NuevoCatalogoServicio(repo)
	usuarioSvc := application.NuevoUsuarioServicio(repo)
	prestamoSvc := application.NuevoPrestamoServicio(repo, repo, repo)
	consultaSvc := application.NuevoConsultaServicio("Biblioteca UIDE", repo, repo, catalogoSvc)

	servicio := application.NuevoGestorBiblioteca(catalogoSvc, usuarioSvc, prestamoSvc, consultaSvc)

	cli.NuevoMenu(servicio).Ejecutar()
}
