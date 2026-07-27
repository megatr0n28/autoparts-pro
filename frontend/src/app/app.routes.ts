import { Routes } from '@angular/router';

import { DashboardComponent } from './dashboard/dashboard';

import { LoginComponent } from './auth/login/login';

import { RegisterComponent } from './auth/register/register';

import {
  MainLayoutComponent,
} from './layout/main-layout/main-layout';

import {
  authGuard,
} from './core/guards/auth.guard';

import {
 VehicleListComponent
} from './vehicles/vehicle-list/vehicle-list';


import {
 VehicleCreateComponent
} from './vehicles/vehicle-create/vehicle-create';

export const routes = [

{

  path: '',

  component: MainLayoutComponent,

  canActivate: [
    authGuard
  ],

  children: [

    {

    path:'vehicles',

    component:
    VehicleListComponent

    },
    {

    path:'vehicles/create',

    component:
    VehicleCreateComponent

    },
    {

      path: '',

      loadComponent: () =>

        import(
          './dashboard/dashboard'
        )
        .then(
          m =>
          m.DashboardComponent
        )

    },

  ]

},



{

 path: 'login',

 loadComponent: () =>

 import(
 './auth/login/login'
 )

 .then(
 m =>
 m.LoginComponent
 )

},



{

 path: 'register',

 loadComponent: () =>

 import(
 './auth/register/register'
 )

 .then(
 m =>
 m.RegisterComponent
 )

}

];