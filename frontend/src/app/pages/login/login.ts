import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

import { AuthService } from '../../core/auth.service';

@Component({
  selector: 'app-login',
  imports: [FormsModule],
  template: `
    <div class="login-wrap">
      <div class="login-card tarjeta">
        <div class="logo">📚</div>
        <h1>Biblioteca UIDE</h1>
        <p class="sub">Sistema de Gestión de Libros</p>

        @if (error()) {
          <div class="alerta error">{{ error() }}</div>
        }

        <form (ngSubmit)="entrar()">
          <label>Usuario</label>
          <input name="username" [(ngModel)]="username" placeholder="admin" autocomplete="username" />

          <label style="margin-top:0.8rem">Contraseña</label>
          <input name="password" type="password" [(ngModel)]="password" placeholder="••••••" autocomplete="current-password" />

          <button type="submit" [disabled]="cargando()" style="width:100%;margin-top:1.2rem">
            {{ cargando() ? 'Ingresando…' : 'Iniciar sesión' }}
          </button>
        </form>

        <p class="hint">Cuenta por defecto: <strong>admin / admin123</strong></p>
      </div>
    </div>
  `,
  styles: [`
    .login-wrap { min-height: 100vh; display: flex; align-items: center; justify-content: center;
      background: linear-gradient(135deg, #1e3a5f, #3182ce); padding: 1rem; }
    .login-card { width: 100%; max-width: 360px; text-align: center; }
    .logo { font-size: 2.6rem; }
    h1 { color: var(--azul); }
    .sub { color: var(--texto-suave); margin-top: 0; }
    form { text-align: left; margin-top: 1rem; }
    .hint { margin-top: 1rem; font-size: 0.78rem; color: var(--texto-suave); }
  `],
})
export class Login {
  private auth = inject(AuthService);
  private router = inject(Router);

  username = 'admin';
  password = 'admin123';
  cargando = signal(false);
  error = signal('');

  entrar(): void {
    this.error.set('');
    this.cargando.set(true);
    this.auth.login(this.username, this.password).subscribe({
      next: () => this.router.navigate(['/']),
      error: (e) => {
        this.error.set(e?.error?.error ?? 'No se pudo iniciar sesión');
        this.cargando.set(false);
      },
    });
  }
}
