import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { BibliotecaService } from '../../core/biblioteca.service';
import { Usuario } from '../../core/models';

@Component({
  selector: 'app-usuarios',
  imports: [FormsModule],
  template: `
    <h2>👥 Usuarios (lectores)</h2>

    @if (mensaje()) { <div class="alerta ok">{{ mensaje() }}</div> }
    @if (error()) { <div class="alerta error">{{ error() }}</div> }

    <div class="grid">
      <div class="tarjeta">
        <h3>Registrar usuario</h3>
        <form (ngSubmit)="crear()">
          <label>Nombre completo</label>
          <input name="n" [(ngModel)]="nombre" required />
          <label style="margin-top:.6rem">Email</label>
          <input name="e" type="email" [(ngModel)]="email" required />
          <button type="submit" style="margin-top:1rem" [disabled]="guardando()">
            {{ guardando() ? 'Guardando…' : 'Registrar' }}
          </button>
        </form>
      </div>

      <div class="tarjeta">
        <table>
          <thead><tr><th>Nombre</th><th>Email</th><th>Estado</th></tr></thead>
          <tbody>
            @for (u of usuarios(); track u.id) {
              <tr>
                <td><strong>{{ u.nombre }}</strong></td>
                <td>{{ u.email }}</td>
                <td>
                  @if (u.activo) { <span class="badge ok">Activo</span> }
                  @else { <span class="badge no">Inactivo</span> }
                </td>
              </tr>
            } @empty {
              <tr><td colspan="3" style="text-align:center;color:var(--texto-suave)">Sin usuarios</td></tr>
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
export class Usuarios {
  private biblioteca = inject(BibliotecaService);

  usuarios = signal<Usuario[]>([]);
  nombre = '';
  email = '';
  guardando = signal(false);
  mensaje = signal('');
  error = signal('');

  constructor() { this.cargar(); }

  cargar(): void {
    this.biblioteca.listarUsuarios().subscribe({
      next: (u) => this.usuarios.set(u),
      error: (e) => this.error.set(e?.error?.error ?? 'Error al cargar'),
    });
  }

  crear(): void {
    this.mensaje.set(''); this.error.set(''); this.guardando.set(true);
    this.biblioteca.crearUsuario({ nombre: this.nombre, email: this.email }).subscribe({
      next: () => {
        this.mensaje.set('Usuario registrado');
        this.nombre = ''; this.email = '';
        this.guardando.set(false);
        this.cargar();
      },
      error: (e) => { this.error.set(e?.error?.error ?? 'No se pudo registrar'); this.guardando.set(false); },
    });
  }
}
