import { Component, inject, signal } from '@angular/core';
import { Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';

import { AuthService } from '../../core/auth.service';
import { BibliotecaService } from '../../core/biblioteca.service';

@Component({
  selector: 'app-layout',
  imports: [RouterOutlet, RouterLink, RouterLinkActive],
  template: `
    <div class="app">
      <aside class="sidebar">
        <div class="brand">📚 <span>Biblioteca UIDE</span></div>
        <nav>
          <a routerLink="/libros" routerLinkActive="activo">📖 Libros</a>
          <a routerLink="/usuarios" routerLinkActive="activo">👥 Usuarios</a>
          <a routerLink="/prestamos" routerLinkActive="activo">🔄 Préstamos</a>
        </nav>
        <div class="pie">
          <small>Arquitectura Hexagonal · Go + SQL Server · JWT</small>
        </div>
      </aside>

      <div class="main">
        <header class="topbar">
          <span class="resumen">{{ resumen() }}</span>
          <div class="usuario">
            <span>👤 {{ auth.usuario() }}</span>
            <button class="secundario" (click)="salir()">Salir</button>
          </div>
        </header>
        <main class="contenido">
          <router-outlet />
        </main>
      </div>
    </div>
  `,
  styles: [`
    .app { display: flex; min-height: 100vh; }
    .sidebar { width: 240px; background: var(--azul); color: #fff; display: flex; flex-direction: column; padding: 1.2rem 1rem; }
    .brand { font-size: 1.1rem; font-weight: 700; margin-bottom: 1.5rem; display: flex; gap: 0.4rem; align-items: center; }
    nav { display: flex; flex-direction: column; gap: 0.3rem; flex: 1; }
    nav a { color: #cbd5e0; text-decoration: none; padding: 0.65rem 0.8rem; border-radius: 6px; font-weight: 600; font-size: 0.9rem; }
    nav a:hover { background: rgba(255,255,255,0.08); color: #fff; }
    nav a.activo { background: var(--acento); color: #fff; }
    .pie { margin-top: auto; color: #7f9cc0; font-size: 0.7rem; line-height: 1.4; }
    .main { flex: 1; display: flex; flex-direction: column; }
    .topbar { background: #fff; border-bottom: 1px solid var(--gris-borde); padding: 0.8rem 1.5rem;
      display: flex; align-items: center; justify-content: space-between; gap: 1rem; }
    .resumen { font-size: 0.8rem; color: var(--texto-suave); }
    .usuario { display: flex; align-items: center; gap: 0.8rem; font-size: 0.85rem; font-weight: 600; }
    .contenido { padding: 1.5rem; flex: 1; }
  `],
})
export class Layout {
  auth = inject(AuthService);
  private biblioteca = inject(BibliotecaService);
  private router = inject(Router);

  resumen = signal('');

  constructor() {
    this.biblioteca.resumen().subscribe({
      next: (r) => this.resumen.set(r.resumen),
      error: () => {},
    });
  }

  salir(): void {
    this.auth.logout();
    this.router.navigate(['/login']);
  }
}
