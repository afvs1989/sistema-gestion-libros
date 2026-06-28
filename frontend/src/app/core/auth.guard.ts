import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

import { TOKEN_KEY } from './api.config';

// authGuard protege las rutas privadas: sin token JWT redirige al login.
export const authGuard: CanActivateFn = () => {
  const router = inject(Router);
  if (localStorage.getItem(TOKEN_KEY)) {
    return true;
  }
  return router.createUrlTree(['/login']);
};
