import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { BibliotecaService } from '../../core/biblioteca.service';
import { CrearLibro, Libro } from '../../core/models';

@Component({
  selector: 'app-libros',
  imports: [FormsModule],
  template: `
    <h2>📖 Catálogo de Libros</h2>

    @if (mensaje()) { <div class="alerta ok">{{ mensaje() }}</div> }
    @if (error()) { <div class="alerta error">{{ error() }}</div> }

    <div class="grid">
      <div class="tarjeta">
        <h3>Registrar libro</h3>
        <form (ngSubmit)="crear()">
          <label>Título</label>
          <input name="t" [(ngModel)]="form.titulo" required />
          <label style="margin-top:.6rem">ISBN</label>
          <input name="i" [(ngModel)]="form.isbn" required />
          <div class="dos">
            <div><label>Autor (nombre)</label><input name="an" [(ngModel)]="form.autorNombre" required /></div>
            <div><label>Autor (apellido)</label><input name="aa" [(ngModel)]="form.autorApellido" required /></div>
          </div>
          <div class="dos">
            <div><label>País</label><input name="ap" [(ngModel)]="form.autorPais" /></div>
            <div><label>Año</label><input name="ay" type="number" [(ngModel)]="form.anio" required /></div>
          </div>
          <label style="margin-top:.6rem">Género</label>
          <input name="g" [(ngModel)]="form.genero" />
          <button type="submit" style="margin-top:1rem" [disabled]="guardando()">
            {{ guardando() ? 'Guardando…' : 'Registrar' }}
          </button>
        </form>
      </div>

      <div class="tarjeta lista">
        <div class="buscador">
          <input placeholder="Buscar por título…" [(ngModel)]="qTitulo" name="qt" />
          <input placeholder="…o por autor" [(ngModel)]="qAutor" name="qa" />
          <button (click)="buscar()">Buscar</button>
          <button class="secundario" (click)="cargar()">Todos</button>
        </div>

        <table>
          <thead>
            <tr><th>Título</th><th>Autor</th><th>Año</th><th>Estado</th><th></th></tr>
          </thead>
          <tbody>
            @for (l of libros(); track l.id) {
              <tr>
                <td><strong>{{ l.titulo }}</strong><br /><small>{{ l.isbn }}</small></td>
                <td>{{ l.autor.nombre }} {{ l.autor.apellido }}</td>
                <td>{{ l.anio }}</td>
                <td>
                  @if (l.disponible) { <span class="badge ok">Disponible</span> }
                  @else { <span class="badge no">Prestado</span> }
                </td>
                <td><button class="peligro" (click)="eliminar(l)">Eliminar</button></td>
              </tr>
            } @empty {
              <tr><td colspan="5" style="text-align:center;color:var(--texto-suave)">Sin libros registrados</td></tr>
            }
          </tbody>
        </table>
      </div>
    </div>
  `,
  styles: [`
    .grid { display: grid; grid-template-columns: 340px 1fr; gap: 1.2rem; align-items: start; }
    .dos { display: grid; grid-template-columns: 1fr 1fr; gap: 0.6rem; margin-top: 0.6rem; }
    .buscador { display: flex; gap: 0.5rem; margin-bottom: 1rem; }
    .buscador input { flex: 1; }
    @media (max-width: 900px) { .grid { grid-template-columns: 1fr; } }
  `],
})
export class Libros {
  private biblioteca = inject(BibliotecaService);

  libros = signal<Libro[]>([]);
  guardando = signal(false);
  mensaje = signal('');
  error = signal('');
  qTitulo = '';
  qAutor = '';

  form: CrearLibro = {
    isbn: '', titulo: '', autorNombre: '', autorApellido: '', autorPais: '', anio: 2024, genero: '',
  };

  constructor() { this.cargar(); }

  cargar(): void {
    this.qTitulo = ''; this.qAutor = '';
    this.biblioteca.listarLibros().subscribe({
      next: (l) => this.libros.set(l),
      error: (e) => this.error.set(e?.error?.error ?? 'Error al cargar'),
    });
  }

  buscar(): void {
    this.biblioteca.buscarLibros(this.qTitulo, this.qAutor).subscribe({
      next: (l) => this.libros.set(l ?? []),
      error: (e) => this.error.set(e?.error?.error ?? 'Error en la búsqueda'),
    });
  }

  crear(): void {
    this.mensaje.set(''); this.error.set(''); this.guardando.set(true);
    this.biblioteca.crearLibro(this.form).subscribe({
      next: () => {
        this.mensaje.set('Libro registrado correctamente');
        this.form = { isbn: '', titulo: '', autorNombre: '', autorApellido: '', autorPais: '', anio: 2024, genero: '' };
        this.guardando.set(false);
        this.cargar();
      },
      error: (e) => { this.error.set(e?.error?.error ?? 'No se pudo registrar'); this.guardando.set(false); },
    });
  }

  eliminar(l: Libro): void {
    this.biblioteca.eliminarLibro(l.id).subscribe({
      next: () => { this.mensaje.set('Libro eliminado'); this.cargar(); },
      error: (e) => this.error.set(e?.error?.error ?? 'No se pudo eliminar'),
    });
  }
}
