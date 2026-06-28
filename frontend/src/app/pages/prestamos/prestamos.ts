import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { DatePipe } from '@angular/common';

import { BibliotecaService } from '../../core/biblioteca.service';
import { Prestamo, Recurso, Usuario } from '../../core/models';

@Component({
  selector: 'app-prestamos',
  imports: [FormsModule, DatePipe],
  template: `
    <h2>🔄 Préstamos</h2>

    @if (mensaje()) { <div class="alerta ok">{{ mensaje() }}</div> }
    @if (error()) { <div class="alerta error">{{ error() }}</div> }

    <div class="grid">
      <div class="tarjeta">
        <h3>Registrar préstamo</h3>
        <form (ngSubmit)="prestar()">
          <label>Recurso disponible</label>
          <select name="r" [(ngModel)]="recursoId" required>
            <option value="">— Selecciona —</option>
            @for (r of recursos(); track r.id) {
              <option [value]="r.id">{{ r.titulo }}</option>
            }
          </select>
          <label style="margin-top:.6rem">Usuario</label>
          <select name="u" [(ngModel)]="usuarioId" required>
            <option value="">— Selecciona —</option>
            @for (u of usuarios(); track u.id) {
              <option [value]="u.id">{{ u.nombre }}</option>
            }
          </select>
          <button type="submit" class="exito" style="margin-top:1rem" [disabled]="guardando()">
            {{ guardando() ? 'Procesando…' : 'Prestar' }}
          </button>
        </form>
      </div>

      <div class="tarjeta">
        <h3>Préstamos activos</h3>
        <table>
          <thead><tr><th>Recurso</th><th>Usuario</th><th>Fecha</th><th></th></tr></thead>
          <tbody>
            @for (p of prestamos(); track p.id) {
              <tr>
                <td>{{ nombreRecurso(p.recursoId) }}</td>
                <td>{{ nombreUsuario(p.usuarioId) }}</td>
                <td>{{ p.fechaPrestamo | date: 'short' }}</td>
                <td><button class="secundario" (click)="devolver(p)">Devolver</button></td>
              </tr>
            } @empty {
              <tr><td colspan="4" style="text-align:center;color:var(--texto-suave)">Sin préstamos activos</td></tr>
            }
          </tbody>
        </table>
      </div>
    </div>
  `,
  styles: [`
    .grid { display: grid; grid-template-columns: 340px 1fr; gap: 1.2rem; align-items: start; }
    @media (max-width: 900px) { .grid { grid-template-columns: 1fr; } }
  `],
})
export class Prestamos {
  private biblioteca = inject(BibliotecaService);

  recursos = signal<Recurso[]>([]);
  usuarios = signal<Usuario[]>([]);
  prestamos = signal<Prestamo[]>([]);
  // Cache de todos los recursos/usuarios para mostrar nombres en la tabla.
  private todosRecursos = signal<Map<string, string>>(new Map());
  private todosUsuarios = signal<Map<string, string>>(new Map());

  recursoId = '';
  usuarioId = '';
  guardando = signal(false);
  mensaje = signal('');
  error = signal('');

  constructor() { this.cargar(); }

  cargar(): void {
    this.biblioteca.recursosDisponibles().subscribe((r) => {
      this.recursos.set(r);
      this.fusionarMapa(this.todosRecursos, r.map((x) => [x.id, x.titulo]));
    });
    this.biblioteca.listarUsuarios().subscribe((u) => {
      this.usuarios.set(u);
      this.fusionarMapa(this.todosUsuarios, u.map((x) => [x.id, x.nombre]));
    });
    this.biblioteca.listarPrestamosActivos().subscribe((p) => this.prestamos.set(p));
  }

  private fusionarMapa(s: ReturnType<typeof signal<Map<string, string>>>, pares: [string, string][]): void {
    const m = new Map(s());
    for (const [k, v] of pares) m.set(k, v);
    s.set(m);
  }

  nombreRecurso(id: string): string { return this.todosRecursos().get(id) ?? id; }
  nombreUsuario(id: string): string { return this.todosUsuarios().get(id) ?? id; }

  prestar(): void {
    this.mensaje.set(''); this.error.set(''); this.guardando.set(true);
    this.biblioteca.prestar(this.recursoId, this.usuarioId).subscribe({
      next: () => {
        this.mensaje.set('Préstamo registrado');
        this.recursoId = ''; this.usuarioId = '';
        this.guardando.set(false);
        this.cargar();
      },
      error: (e) => { this.error.set(e?.error?.error ?? 'No se pudo prestar'); this.guardando.set(false); },
    });
  }

  devolver(p: Prestamo): void {
    this.biblioteca.devolver(p.recursoId).subscribe({
      next: () => { this.mensaje.set('Recurso devuelto'); this.cargar(); },
      error: (e) => this.error.set(e?.error?.error ?? 'No se pudo devolver'),
    });
  }
}
