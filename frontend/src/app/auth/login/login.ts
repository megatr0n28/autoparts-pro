import {
  Component,
} from '@angular/core';

import {
  FormBuilder,
  ReactiveFormsModule,
  Validators,
  FormGroup,
} from '@angular/forms';

import {
  MatInputModule,
} from '@angular/material/input';

import {
  MatButtonModule,
} from '@angular/material/button';

import {
  MatCardModule,
} from '@angular/material/card';


@Component({

  selector: 'app-login',

  standalone: true,

  imports: [

    ReactiveFormsModule,

    MatInputModule,

    MatButtonModule,

    MatCardModule,

  ],

  templateUrl: './login.html',

  styleUrl: './login.scss',

})

export class LoginComponent {


  loginForm: FormGroup;


  constructor(
    private fb: FormBuilder,
  ) {


    this.loginForm =
      this.fb.group({

        email: [

          '',

          [

            Validators.required,

            Validators.email,

          ],

        ],


        password: [

          '',

          [

            Validators.required,

          ],

        ],


      });


  }



  submit() {


    if (
      this.loginForm.invalid
    ) {

      return;

    }


    console.log(
      this.loginForm.value
    );


  }


}