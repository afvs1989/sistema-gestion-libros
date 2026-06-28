import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';

import { TOKEN_KEY } from './api.config';

// authInterceptor añade el header Authorization: Bearer <token> a cada petición
// y redirige al login si la API responde 401 (token expirado o ausente).
export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const router = inject(Router);
  const token = localStorage.getItem(TOKEN_KEY);

  const peticion = token
    ? req.clone({ setHeaders: { Authorization: `Bearer ${token}` } })
    : req;

  return next(peticion).pipe(
    catchError((err) => {
      if (err.status === 401) {
        localStorage.clear();
        router.navigate(['/login']);
      }
      return throwError(() => err);
    }),
  );
};
