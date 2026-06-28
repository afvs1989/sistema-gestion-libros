import { Routes } from '@angular/router';

import { authGuard } from './core/auth.guard';
import { Login } from './pages/login/login';
import { Layout } from './pages/layout/layout';
import { Libros } from './pages/libros/libros';
import { Usuarios } from './pages/usuarios/usuarios';
import { Prestamos } from './pages/prestamos/prestamos';

export const routes: Routes = [
  { path: 'login', component: Login },
  {
    path: '',
    component: Layout,
    canActivate: [authGuard],
    children: [
      { path: '', redirectTo: 'libros', pathMatch: 'full' },
      { path: 'libros', component: Libros },
      { path: 'usuarios', component: Usuarios },
      { path: 'prestamos', component: Prestamos },
    ],
  },
  { path: '**', redirectTo: '' },
];
