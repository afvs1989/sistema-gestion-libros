package entity

// Recurso define el contrato común para elementos catalogables de la biblioteca.
// Principio ISP: interfaz mínima para operaciones polimórficas de préstamo y consulta.
type Recurso interface {
	ID() string
	Titulo() string
	Disponible() bool
	MarcarPrestado()
	MarcarDevuelto()
	Descripcion() string
}
