package entity

import "time"

// Este archivo contiene las funciones de "rehidratación" usadas por los adaptadores
// de persistencia (patrón hexagonal): permiten reconstruir entidades de dominio a
// partir de datos ya almacenados en la base de datos, sin volver a ejecutar las
// validaciones de los constructores (los datos ya fueron validados al crearse).
// Al pertenecer al mismo paquete pueden escribir los campos privados sin exponer setters.

// ReconstruirAutor rehidrata un Autor desde la persistencia.
func ReconstruirAutor(nombre, apellido, pais string) *Autor {
	return &Autor{nombre: nombre, apellido: apellido, pais: pais}
}

// Getters adicionales del Autor necesarios para el mapeo a modelos de persistencia.
func (a *Autor) Nombre() string   { return a.nombre }
func (a *Autor) Apellido() string { return a.apellido }

// ReconstruirLibro rehidrata un Libro desde la persistencia.
func ReconstruirLibro(id, isbn, titulo string, autor *Autor, anio int, genero string, disponible bool, fechaRegistro time.Time) *Libro {
	return &Libro{
		id:            id,
		isbn:          isbn,
		titulo:        titulo,
		autor:         autor,
		anio:          anio,
		genero:        genero,
		disponible:    disponible,
		fechaRegistro: fechaRegistro,
	}
}

// ReconstruirRevista rehidrata una Revista desde la persistencia.
func ReconstruirRevista(id, titulo, editorial string, numero int, disponible bool) *Revista {
	return &Revista{
		id:         id,
		titulo:     titulo,
		numero:     numero,
		editorial:  editorial,
		disponible: disponible,
	}
}

// ReconstruirUsuario rehidrata un Usuario desde la persistencia.
func ReconstruirUsuario(id, nombre, email string, activo bool) *Usuario {
	return &Usuario{id: id, nombre: nombre, email: email, activo: activo}
}

// ReconstruirPrestamo rehidrata un Prestamo desde la persistencia.
func ReconstruirPrestamo(id, usuarioID, recursoID string, fechaPrestamo time.Time, fechaDevolucion *time.Time, activo bool) *Prestamo {
	return &Prestamo{
		id:              id,
		usuarioID:       usuarioID,
		recursoID:       recursoID,
		fechaPrestamo:   fechaPrestamo,
		fechaDevolucion: fechaDevolucion,
		activo:          activo,
	}
}

// FechaDevolucion expone la fecha de devolución (nil si el préstamo sigue activo).
func (p *Prestamo) FechaDevolucion() *time.Time { return p.fechaDevolucion }
