package persistence

import (
	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/port/outbound"
)

// Memoria es el adaptador outbound que persiste datos en mapas en RAM.
type Memoria struct {
	libros           map[string]*entity.Libro
	isbnIndex        map[string]string
	recursos         map[string]entity.Recurso
	usuarios         map[string]*entity.Usuario
	prestamos        map[string]*entity.Prestamo
	prestamosActivos map[string]string
}

var _ outbound.Catalogo = (*Memoria)(nil)
var _ outbound.Usuarios = (*Memoria)(nil)
var _ outbound.Prestamos = (*Memoria)(nil)

func NuevaMemoria() *Memoria {
	return &Memoria{
		libros:           make(map[string]*entity.Libro),
		isbnIndex:        make(map[string]string),
		recursos:         make(map[string]entity.Recurso),
		usuarios:         make(map[string]*entity.Usuario),
		prestamos:        make(map[string]*entity.Prestamo),
		prestamosActivos: make(map[string]string),
	}
}
