package cli

import "fmt"

func (m *Menu) prestarRecurso() {
	recursoID := m.leerLinea("ID del recurso: ")
	usuarioID := m.leerLinea("ID del usuario: ")
	prestamoID := m.leerLinea("ID del préstamo: ")

	prestamo, err := m.svc.PrestarRecurso(recursoID, usuarioID, prestamoID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Préstamo registrado:", prestamo.Resumen())
}

func (m *Menu) devolverRecurso() {
	recursoID := m.leerLinea("ID del recurso a devolver: ")
	if err := m.svc.DevolverRecurso(recursoID); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Recurso devuelto correctamente.")
}

func (m *Menu) listarPrestamosActivos() {
	prestamos := m.svc.ListarPrestamosActivos()
	if len(prestamos) == 0 {
		fmt.Println("No hay préstamos activos.")
		return
	}
	for _, p := range prestamos {
		fmt.Println(" •", p.Resumen())
	}
}

func (m *Menu) cargarDemo() {
	if err := CargarDatosDemo(m.svc); err != nil {
		fmt.Println("Error al cargar demo:", err)
		return
	}
	fmt.Println("Datos de demostración cargados correctamente.")
	fmt.Println(m.svc.ResumenCatalogo())
}
