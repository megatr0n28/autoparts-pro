import { Routes } from '@angular/router';

import { MainLayoutComponent } from './layout/main-layout/main-layout';

import { authGuard } from './core/guards/auth.guard';

import { VehicleListComponent } from './vehicles/vehicle-list/vehicle-list';

import { VehicleCreateComponent } from './vehicles/vehicle-create/vehicle-create';

import { PartsSearchComponent } from './parts/parts-search/parts-search';

export const routes: Routes = [
  {
    path: '',
    component: MainLayoutComponent,
    canActivate: [authGuard],

    children: [
      {
        path: 'vehicles/create',
        component: VehicleCreateComponent,
      },

      {
        path: 'vehicles',
        component: VehicleListComponent,
      },

      {
        path: 'parts',
        component: PartsSearchComponent,
      },

      {
        path: '',
        loadComponent: () =>
          import('./dashboard/dashboard')
            .then(m => m.DashboardComponent),
      },
    ],
  },

  {
    path: 'login',
    loadComponent: () =>
      import('./auth/login/login')
        .then(m => m.LoginComponent),
  },

  {
    path: 'register',
    loadComponent: () =>
      import('./auth/register/register')
        .then(m => m.RegisterComponent),
  },
];