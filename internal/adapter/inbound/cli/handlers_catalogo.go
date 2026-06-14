package cli

import (
	"fmt"
	"strconv"
	"strings"

	"sistema-gestion-libros/internal/domain/entity"
)

func (m *Menu) registrarLibro() {
	id := m.leerLinea("ID del libro: ")
	isbn := m.leerLinea("ISBN: ")
	titulo := m.leerLinea("Título: ")
	nombreAutor := m.leerLinea("Nombre del autor: ")
	apellidoAutor := m.leerLinea("Apellido del autor: ")
	pais := m.leerLinea("País del autor: ")
	anio := m.leerEntero("Año de publicación: ")
	genero := m.leerLinea("Género: ")

	autor, err := entity.NuevoAutor(nombreAutor, apellidoAutor, pais)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	libro, err := entity.NuevoLibro(id, isbn, titulo, autor, anio, genero)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := m.svc.RegistrarLibro(libro); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Libro registrado:", libro.Descripcion())
}

func (m *Menu) registrarRevista() {
	id := m.leerLinea("ID de la revista: ")
	titulo := m.leerLinea("Título: ")
	editorial := m.leerLinea("Editorial: ")
	numero := m.leerEntero("Número de edición: ")

	revista, err := entity.NuevaRevista(id, titulo, editorial, numero)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := m.svc.RegistrarRecurso(revista); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Revista registrada:", revista.Descripcion())
}

func (m *Menu) buscarPorTitulo() {
	termino := m.leerLinea("Término de búsqueda: ")
	m.imprimirLibros(m.svc.BuscarLibrosPorTitulo(termino))
}

func (m *Menu) buscarPorAutor() {
	autor := m.leerLinea("Nombre del autor: ")
	m.imprimirLibros(m.svc.BuscarLibrosPorAutor(autor))
}

func (m *Menu) listarLibros() {
	m.imprimirLibros(m.svc.ListarLibros())
}

func (m *Menu) listarDisponibles() {
	recursos := m.svc.ListarRecursosDisponibles()
	if len(recursos) == 0 {
		fmt.Println("No hay recursos disponibles.")
		return
	}
	for _, r := range recursos {
		fmt.Println(" •", r.Descripcion())
	}
}

func (m *Menu) imprimirLibros(libros []*entity.Libro) {
	if len(libros) == 0 {
		fmt.Println("No se encontraron libros.")
		return
	}
	for _, libro := range libros {
		fmt.Println(" •", libro.Descripcion())
	}
}

func (m *Menu) leerLinea(prompt string) string {
	fmt.Print(prompt)
	texto, _ := m.reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func (m *Menu) leerEntero(prompt string) int {
	for {
		texto := m.leerLinea(prompt)
		valor, err := strconv.Atoi(texto)
		if err == nil {
			return valor
		}
		fmt.Println("Ingrese un número válido.")
	}
}
