import {
  Component,
} from '@angular/core';


import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
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

  selector: 'app-register',

  standalone: true,

  imports: [

    ReactiveFormsModule,

    MatInputModule,

    MatButtonModule,

    MatCardModule,

  ],

  templateUrl: './register.html',

  styleUrl: './register.scss',

})


export class RegisterComponent {


  registerForm: FormGroup;


  constructor(
    private fb: FormBuilder,
  ) {


    this.registerForm =
      this.fb.group({

        firstName: [

          '',

          [

            Validators.required,

          ],

        ],


        lastName: [

          '',

          [

            Validators.required,

          ],

        ],


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

            Validators.minLength(8),

          ],

        ],


      });


  }



  submit() {


    if (
      this.registerForm.invalid
    ) {

      return;

    }


    console.log(
      "Register submitted",
      this.registerForm.value
    );


  }


}