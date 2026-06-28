package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"sistema-gestion-libros/internal/port/inbound"
)

// Menu es el adaptador inbound que expone los casos de uso mediante consola interactiva.
type Menu struct {
	svc    inbound.Biblioteca
	reader *bufio.Reader
}

func NuevoMenu(svc inbound.Biblioteca) *Menu {
	return &Menu{
		svc:    svc,
		reader: bufio.NewReader(os.Stdin),
	}
}

// Ejecutar inicia el bucle principal del menú hasta que el usuario elija salir.
func (m *Menu) Ejecutar() {
	for {
		m.mostrarMenu()
		opcion := strings.TrimSpace(m.leerLinea("Seleccione una opción: "))
		if opcion == "0" {
			fmt.Println("\n¡Hasta pronto!")
			return
		}
		m.procesarOpcion(opcion)
		fmt.Println("\nPresione Enter para continuar...")
		_, _ = m.reader.ReadString('\n')
	}
}

func (m *Menu) mostrarMenu() {
	fmt.Println("\n========================================")
	fmt.Printf("  %s\n", m.svc.Nombre())
	fmt.Println("========================================")
	fmt.Println(" 1. Registrar libro")
	fmt.Println(" 2. Registrar revista")
	fmt.Println(" 3. Registrar usuario")
	fmt.Println(" 4. Buscar libros por título")
	fmt.Println(" 5. Buscar libros por autor")
	fmt.Println(" 6. Listar todos los libros")
	fmt.Println(" 7. Listar recursos disponibles")
	fmt.Println(" 8. Listar usuarios")
	fmt.Println(" 9. Prestar recurso")
	fmt.Println("10. Devolver recurso")
	fmt.Println("11. Ver préstamos activos")
	fmt.Println("12. Resumen del catálogo")
	fmt.Println("13. Cargar datos de demostración")
	fmt.Println(" 0. Salir")
	fmt.Println("----------------------------------------")
}

func (m *Menu) procesarOpcion(opcion string) {
	switch opcion {
	case "1":
		m.registrarLibro()
	case "2":
		m.registrarRevista()
	case "3":
		m.registrarUsuario()
	case "4":
		m.buscarPorTitulo()
	case "5":
		m.buscarPorAutor()
	case "6":
		m.listarLibros()
	case "7":
		m.listarDisponibles()
	case "8":
		m.listarUsuarios()
	case "9":
		m.prestarRecurso()
	case "10":
		m.devolverRecurso()
	case "11":
		m.listarPrestamosActivos()
	case "12":
		fmt.Println("\n>>", m.svc.ResumenCatalogo())
	case "13":
		m.cargarDemo()
	default:
		fmt.Println("Opción no válida.")
	}
}
