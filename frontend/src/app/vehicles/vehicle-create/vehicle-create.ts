import {
  Component,
  OnDestroy,
  ChangeDetectorRef,
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
  CommonModule,
} from '@angular/common';


import {
  VehicleService,
} from '../../core/services/vehicle.service';



@Component({

  selector: 'app-vehicle-create',

  standalone: true,

  imports: [

    CommonModule,

    ReactiveFormsModule,

    MatInputModule,

    MatButtonModule,

    MatCardModule,

  ],

  templateUrl:
    './vehicle-create.html',

  styleUrl:
    './vehicle-create.scss',

})


export class VehicleCreateComponent implements OnDestroy{


  vehicleForm: FormGroup;


  submitting = false;


  error = '';



  constructor(

    private fb: FormBuilder,

    private vehicleService: VehicleService,

    private router: Router,
    private cd: ChangeDetectorRef,

  ) {


    this.vehicleForm =

      this.fb.group({

        year: [

          '',

          [

            Validators.required,

            Validators.min(1900),

            Validators.max(
              new Date().getFullYear()
            ),

          ],

        ],


        make: [

          '',

          [

            Validators.required,

            Validators.minLength(2),

          ],

        ],


        model: [

          '',

          [

            Validators.required,

            Validators.minLength(1),

          ],

        ],


        vin: [

          '',

        ],


        license_plate: [

          '',

        ],

      });


  }



  submit(): void {


    /*
      Clear previous errors
    */

    this.error = '';



    if (
      this.vehicleForm.invalid
    ) {


      this.vehicleForm.markAllAsTouched();


      return;

    }



    this.submitting = true;



    /*
      Angular form values return strings.
      Go expects year as int.
    */

    const vehicle = {

      ...this.vehicleForm.value,


      year:

        Number(
          this.vehicleForm.value.year
        ),

    };



    this.vehicleService

      .createVehicle(
        vehicle
      )

      .subscribe({

        next: () => {

          this.submitting = false;


          this.router.navigate(
            ['/vehicles'],
            {
              queryParams: {
                message: 'Vehicle created successfully'
              }
            }
          );

        },


        error: (err) => {

          console.log(
            "HTTP ERROR RECEIVED",
            err
          );


          this.error =
            err.error?.error ??
            "Unable to create vehicle";


          this.submitting = false;


          this.cd.detectChanges();

        },

      });


  }
  ngOnDestroy(): void {

  console.log(
    "VehicleCreateComponent destroyed"
  );

}


}