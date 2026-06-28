// Modelos que reflejan los DTOs JSON expuestos por la API en Go.

export interface Autor {
  nombre: string;
  apellido: string;
  pais: string;
}

export interface Libro {
  id: string;
  isbn: string;
  titulo: string;
  autor: Autor;
  anio: number;
  genero: string;
  disponible: boolean;
  fechaRegistro: string;
}

export interface Recurso {
  id: string;
  titulo: string;
  disponible: boolean;
  descripcion: string;
}

export interface Usuario {
  id: string;
  nombre: string;
  email: string;
  activo: boolean;
}

export interface Prestamo {
  id: string;
  usuarioId: string;
  recursoId: string;
  fechaPrestamo: string;
  fechaDevolucion: string | null;
  activo: boolean;
}

export interface TokenResponse {
  token: string;
  tipo: string;
  username: string;
  rol: string;
  expira: string;
}

export interface CrearLibro {
  isbn: string;
  titulo: string;
  autorNombre: string;
  autorApellido: string;
  autorPais: string;
  anio: number;
  genero: string;
}

export interface CrearRevista {
  titulo: string;
  editorial: string;
  numero: number;
}

export interface CrearUsuario {
  nombre: string;
  email: string;
}
