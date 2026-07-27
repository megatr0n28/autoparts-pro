import { Routes } from '@angular/router';

import { DashboardComponent } from './dashboard/dashboard';

import { LoginComponent } from './auth/login/login';

import { RegisterComponent } from './auth/register/register';

import {
  authGuard,
} from './core/guards/auth.guard';

export const routes: Routes = [

  {

    path: '',

    component: DashboardComponent,
    canActivate: [
      authGuard
    ],

  },

  {

    path: 'login',

    component: LoginComponent,

  },

  {

    path: 'register',

    component: RegisterComponent,

  },

  {

    path: '**',

    redirectTo: '',

  },

];