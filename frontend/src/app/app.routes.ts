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
        path: 'vehicles/:id/edit',
        loadComponent: () =>
          import('./vehicles/vehicle-edit/vehicle-edit')
            .then(m => m.VehicleEditComponent),
      },

      {
        path: 'vehicles',
        component: VehicleListComponent,
      },

      {
        path: 'parts/search',
        component: PartsSearchComponent,
      },

      {
        path: 'dashboard',
        loadComponent: () =>
          import('./dashboard/dashboard')
            .then(m => m.DashboardComponent),
      },

      {
        path: 'admin',
        loadComponent: () =>
          import('./admin/admin')
            .then(m => m.AdminComponent),
      },

      {
        path: 'admin/vehicles/:id/edit',
        loadComponent: () =>
          import('./admin/admin-vehicle-edit')
            .then(m => m.AdminVehicleEditComponent),
      },

      {
        path: 'admin/customers/:id/edit',
        loadComponent: () =>
          import('./admin/admin-customer-edit')
            .then(m => m.AdminCustomerEditComponent),
      },

      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'dashboard',
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