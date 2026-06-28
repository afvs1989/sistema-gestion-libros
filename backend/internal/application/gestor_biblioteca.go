package application

import (
	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/port/inbound"
)

// GestorBiblioteca es la fachada que compone servicios especializados (DIP + OCP).
type GestorBiblioteca struct {
	catalogo  inbound.CatalogoCasosUso
	usuarios  inbound.UsuarioCasosUso
	prestamos inbound.PrestamoCasosUso
	consulta  inbound.ConsultaBiblioteca
}

var _ inbound.Biblioteca = (*GestorBiblioteca)(nil)

func NuevoGestorBiblioteca(
	catalogo inbound.CatalogoCasosUso,
	usuarios inbound.UsuarioCasosUso,
	prestamos inbound.PrestamoCasosUso,
	consulta inbound.ConsultaBiblioteca,
) *GestorBiblioteca {
	return &GestorBiblioteca{
		catalogo:  catalogo,
		usuarios:  usuarios,
		prestamos: prestamos,
		consulta:  consulta,
	}
}

func (g *GestorBiblioteca) Nombre() string        { return g.consulta.Nombre() }
func (g *GestorBiblioteca) ResumenCatalogo() string { return g.consulta.ResumenCatalogo() }

func (g *GestorBiblioteca) RegistrarLibro(libro *entity.Libro) error {
	return g.catalogo.RegistrarLibro(libro)
}

func (g *GestorBiblioteca) RegistrarRecurso(recurso entity.Recurso) error {
	return g.catalogo.RegistrarRecurso(recurso)
}

func (g *GestorBiblioteca) BuscarLibrosPorTitulo(termino string) []*entity.Libro {
	return g.catalogo.BuscarLibrosPorTitulo(termino)
}

func (g *GestorBiblioteca) BuscarLibrosPorAutor(nombreAutor string) []*entity.Libro {
	return g.catalogo.BuscarLibrosPorAutor(nombreAutor)
}

func (g *GestorBiblioteca) ListarLibros() []*entity.Libro {
	return g.catalogo.ListarLibros()
}

func (g *GestorBiblioteca) ListarRecursosDisponibles() []entity.Recurso {
	return g.catalogo.ListarRecursosDisponibles()
}

func (g *GestorBiblioteca) EliminarLibro(id string) error {
	return g.catalogo.EliminarLibro(id)
}

func (g *GestorBiblioteca) RegistrarUsuario(usuario *entity.Usuario) error {
	return g.usuarios.RegistrarUsuario(usuario)
}

func (g *GestorBiblioteca) ListarUsuarios() []*entity.Usuario {
	return g.usuarios.ListarUsuarios()
}

func (g *GestorBiblioteca) PrestarRecurso(recursoID, usuarioID, prestamoID string) (*entity.Prestamo, error) {
	return g.prestamos.PrestarRecurso(recursoID, usuarioID, prestamoID)
}

func (g *GestorBiblioteca) DevolverRecurso(recursoID string) error {
	return g.prestamos.DevolverRecurso(recursoID)
}

func (g *GestorBiblioteca) ListarPrestamosActivos() []*entity.Prestamo {
	return g.prestamos.ListarPrestamosActivos()
}
