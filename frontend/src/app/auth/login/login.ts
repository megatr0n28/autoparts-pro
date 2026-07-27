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
  Router,
} from '@angular/router';


import {
  MatInputModule,
} from '@angular/material/input';


import {
  MatButtonModule,
} from '@angular/material/button';


import {
  MatCardModule,
} from '@angular/material/card';


import {
  AuthService,
} from '../../core/services/auth.service';



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


  errorMessage = '';



  constructor(

    private fb: FormBuilder,

    private authService: AuthService,

    private router: Router,

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




  submit(): void {


    this.errorMessage = '';



    if (
      this.loginForm.invalid
    ) {

      return;

    }



    const {

      email,

      password,

    } = this.loginForm.value;



    this.authService.login({

      email,

      password,

    })

    .subscribe({

      next: () => {


        this.router.navigate([
          '/'
        ]);


      },


      error: (error) => {


        console.error(
          "Login failed",
          error
        );


        this.errorMessage =
          "Invalid email or password";


      },


    });


  }


}