package inbound

// ConsultaBiblioteca expone consultas generales del sistema.
type ConsultaBiblioteca interface {
	Nombre() string
	ResumenCatalogo() string
}

// Biblioteca compone los puertos inbound especializados (facade para adaptadores).
type Biblioteca interface {
	CatalogoCasosUso
	UsuarioCasosUso
	PrestamoCasosUso
	ConsultaBiblioteca
}
