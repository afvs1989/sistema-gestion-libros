package cli

import (
	"fmt"

	"sistema-gestion-libros/internal/domain/entity"
)

func (m *Menu) registrarUsuario() {
	id := m.leerLinea("ID del usuario: ")
	nombre := m.leerLinea("Nombre: ")
	email := m.leerLinea("Email: ")

	usuario, err := entity.NuevoUsuario(id, nombre, email)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := m.svc.RegistrarUsuario(usuario); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Usuario registrado: %s (%s)\n", usuario.Nombre(), usuario.Email())
}

func (m *Menu) listarUsuarios() {
	usuarios := m.svc.ListarUsuarios()
	if len(usuarios) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}
	for _, u := range usuarios {
		fmt.Printf(" • [%s] %s — %s\n", u.ID(), u.Nombre(), u.Email())
	}
}
