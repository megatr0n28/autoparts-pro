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


  errorMessage = '';

  successMessage = '';



  constructor(

    private fb: FormBuilder,

    private authService: AuthService,

    private router: Router,

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




  submit(): void {


    this.errorMessage = '';

    this.successMessage = '';



    if (
      this.registerForm.invalid
    ) {

      return;

    }



    const form =
      this.registerForm.value;



    const request = {


      first_name:
        form.firstName,


      last_name:
        form.lastName,


      email:
        form.email,


      password:
        form.password,


    };



    this.authService.register(
      request
    )

    .subscribe({


      next: () => {


        this.successMessage =
          "Registration successful. Redirecting to login...";



        setTimeout(() => {


          this.router.navigate([
            '/login'
          ]);


        }, 1500);



      },


      error: (error) => {


        console.error(
          "Registration failed",
          error
        );


        this.errorMessage =
          error.error?.error ??
          "Registration failed";


      },


    });


  }


}