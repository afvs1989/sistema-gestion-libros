package sqlserver

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
	"sistema-gestion-libros/internal/port/outbound"
)

// RepoCatalogo implementa el puerto outbound.Catalogo sobre SQL Server.
type RepoCatalogo struct {
	db *gorm.DB
}

var _ outbound.Catalogo = (*RepoCatalogo)(nil)

func NuevoRepoCatalogo(db *gorm.DB) *RepoCatalogo { return &RepoCatalogo{db: db} }

// ---- Mapeo modelo <-> entidad ----

func libroAModelo(l *entity.Libro) LibroModel {
	return LibroModel{
		ID:            l.ID(),
		ISBN:          l.ISBN(),
		Titulo:        l.Titulo(),
		AutorNombre:   l.Autor().Nombre(),
		AutorApellido: l.Autor().Apellido(),
		AutorPais:     l.Autor().Pais(),
		Anio:          l.Anio(),
		Genero:        l.Genero(),
		Disponible:    l.Disponible(),
		FechaRegistro: l.FechaRegistro(),
	}
}

func modeloALibro(m LibroModel) *entity.Libro {
	autor := entity.ReconstruirAutor(m.AutorNombre, m.AutorApellido, m.AutorPais)
	return entity.ReconstruirLibro(m.ID, m.ISBN, m.Titulo, autor, m.Anio, m.Genero, m.Disponible, m.FechaRegistro)
}

func revistaAModelo(r *entity.Revista) RevistaModel {
	return RevistaModel{
		ID:         r.ID(),
		Titulo:     r.Titulo(),
		Numero:     r.Numero(),
		Editorial:  r.Editorial(),
		Disponible: r.Disponible(),
	}
}

func modeloARevista(m RevistaModel) *entity.Revista {
	return entity.ReconstruirRevista(m.ID, m.Titulo, m.Editorial, m.Numero, m.Disponible)
}

// ---- Libros ----

func (r *RepoCatalogo) GuardarLibro(libro *entity.Libro) error {
	var existente LibroModel
	err := r.db.Where("isbn = ?", libro.ISBN()).First(&existente).Error
	if err == nil && existente.ID != libro.ID() {
		return errores.ErrISBNDuplicado
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("verificar ISBN: %w", err)
	}
	if err := r.db.Save(libroAModelo(libro)).Error; err != nil {
		return fmt.Errorf("guardar libro: %w", err)
	}
	return nil
}

func (r *RepoCatalogo) ObtenerLibroPorID(id string) (*entity.Libro, error) {
	var m LibroModel
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errores.ErrLibroNoEncontrado
		}
		return nil, err
	}
	return modeloALibro(m), nil
}

func (r *RepoCatalogo) ObtenerLibroPorISBN(isbn string) (*entity.Libro, error) {
	var m LibroModel
	if err := r.db.First(&m, "isbn = ?", isbn).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errores.ErrLibroNoEncontrado
		}
		return nil, err
	}
	return modeloALibro(m), nil
}

func (r *RepoCatalogo) ListarLibros() []*entity.Libro {
	var modelos []LibroModel
	r.db.Order("fecha_registro asc").Find(&modelos)
	libros := make([]*entity.Libro, 0, len(modelos))
	for _, m := range modelos {
		libros = append(libros, modeloALibro(m))
	}
	return libros
}

func (r *RepoCatalogo) EliminarLibro(id string) error {
	res := r.db.Delete(&LibroModel{}, "id = ?", id)
	if res.Error != nil {
		return fmt.Errorf("eliminar libro: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return errores.ErrLibroNoEncontrado
	}
	return nil
}

// ---- Recursos (polimorfismo Libro/Revista) ----

func (r *RepoCatalogo) GuardarRecurso(recurso entity.Recurso) error {
	switch v := recurso.(type) {
	case *entity.Libro:
		if err := r.db.Save(libroAModelo(v)).Error; err != nil {
			return fmt.Errorf("guardar recurso (libro): %w", err)
		}
	case *entity.Revista:
		if err := r.db.Save(revistaAModelo(v)).Error; err != nil {
			return fmt.Errorf("guardar recurso (revista): %w", err)
		}
	default:
		return fmt.Errorf("%w: tipo de recurso no soportado", errores.ErrCampoInvalido)
	}
	return nil
}

func (r *RepoCatalogo) ObtenerRecursoPorID(id string) (entity.Recurso, error) {
	var libro LibroModel
	if err := r.db.First(&libro, "id = ?", id).Error; err == nil {
		return modeloALibro(libro), nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var revista RevistaModel
	if err := r.db.First(&revista, "id = ?", id).Error; err == nil {
		return modeloARevista(revista), nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return nil, fmt.Errorf("%w: recurso con id %s", errores.ErrLibroNoEncontrado, id)
}

func (r *RepoCatalogo) ListarRecursos() []entity.Recurso {
	var recursos []entity.Recurso

	var libros []LibroModel
	r.db.Find(&libros)
	for _, m := range libros {
		recursos = append(recursos, modeloALibro(m))
	}

	var revistas []RevistaModel
	r.db.Find(&revistas)
	for _, m := range revistas {
		recursos = append(recursos, modeloARevista(m))
	}
	return recursos
}
