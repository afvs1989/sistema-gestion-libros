import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { API_URL } from './api.config';
import {
  CrearLibro,
  CrearRevista,
  CrearUsuario,
  Libro,
  Prestamo,
  Recurso,
  Usuario,
} from './models';

// BibliotecaService consume los servicios web REST del backend (catálogo,
// usuarios, préstamos y resumen). El token JWT lo añade el interceptor.
@Injectable({ providedIn: 'root' })
export class BibliotecaService {
  private http = inject(HttpClient);

  // ---- Libros ----
  listarLibros(): Observable<Libro[]> {
    return this.http.get<Libro[]>(`${API_URL}/libros`);
  }
  crearLibro(libro: CrearLibro): Observable<Libro> {
    return this.http.post<Libro>(`${API_URL}/libros`, libro);
  }
  buscarLibros(titulo: string, autor: string): Observable<Libro[]> {
    const params: Record<string, string> = {};
    if (autor) params['autor'] = autor;
    else if (titulo) params['titulo'] = titulo;
    return this.http.get<Libro[]>(`${API_URL}/libros/buscar`, { params });
  }
  eliminarLibro(id: string): Observable<unknown> {
    return this.http.delete(`${API_URL}/libros/${id}`);
  }

  // ---- Revistas y recursos ----
  crearRevista(revista: CrearRevista): Observable<Recurso> {
    return this.http.post<Recurso>(`${API_URL}/revistas`, revista);
  }
  recursosDisponibles(): Observable<Recurso[]> {
    return this.http.get<Recurso[]>(`${API_URL}/recursos/disponibles`);
  }

  // ---- Usuarios ----
  listarUsuarios(): Observable<Usuario[]> {
    return this.http.get<Usuario[]>(`${API_URL}/usuarios`);
  }
  crearUsuario(usuario: CrearUsuario): Observable<Usuario> {
    return this.http.post<Usuario>(`${API_URL}/usuarios`, usuario);
  }

  // ---- Préstamos ----
  listarPrestamosActivos(): Observable<Prestamo[]> {
    return this.http.get<Prestamo[]>(`${API_URL}/prestamos/activos`);
  }
  prestar(recursoId: string, usuarioId: string): Observable<Prestamo> {
    return this.http.post<Prestamo>(`${API_URL}/prestamos`, { recursoId, usuarioId });
  }
  devolver(recursoId: string): Observable<unknown> {
    return this.http.put(`${API_URL}/prestamos/devolver/${recursoId}`, {});
  }

  // ---- Resumen ----
  resumen(): Observable<{ biblioteca: string; resumen: string }> {
    return this.http.get<{ biblioteca: string; resumen: string }>(`${API_URL}/catalogo/resumen`);
  }
}
