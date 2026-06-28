import { Injectable, computed, inject, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';

import { API_URL, TOKEN_KEY, USER_KEY } from './api.config';
import { TokenResponse } from './models';

// AuthService gestiona el login contra la API, el almacenamiento del token JWT
// y el estado de sesión reactivo (signals).
@Injectable({ providedIn: 'root' })
export class AuthService {
  private http = inject(HttpClient);

  private readonly _usuario = signal<string | null>(localStorage.getItem(USER_KEY));
  readonly usuario = this._usuario.asReadonly();
  readonly autenticado = computed(() => this._usuario() !== null);

  login(username: string, password: string): Observable<TokenResponse> {
    return this.http.post<TokenResponse>(`${API_URL}/auth/login`, { username, password }).pipe(
      tap((res) => {
        localStorage.setItem(TOKEN_KEY, res.token);
        localStorage.setItem(USER_KEY, res.username);
        this._usuario.set(res.username);
      }),
    );
  }

  registrar(username: string, password: string, rol: string): Observable<unknown> {
    return this.http.post(`${API_URL}/auth/register`, { username, password, rol });
  }

  logout(): void {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
    this._usuario.set(null);
  }

  get token(): string | null {
    return localStorage.getItem(TOKEN_KEY);
  }
}
