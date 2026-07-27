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
  VehicleService,
} from '../../core/services/vehicle.service';



@Component({

  selector:'app-vehicle-create',

  standalone:true,

  imports:[

    ReactiveFormsModule,

    MatInputModule,

    MatButtonModule,

  ],

  templateUrl:
    './vehicle-create.html',

  styleUrl:
    './vehicle-create.scss',

})


export class VehicleCreateComponent {


  vehicleForm: FormGroup;



  constructor(

    private fb: FormBuilder,

    private service: VehicleService,

    private router: Router,

  ){


    this.vehicleForm =
      this.fb.group({

        year:[

          '',

          Validators.required

        ],


        make:[

          '',

          Validators.required

        ],


        model:[

          '',

          Validators.required

        ],


        vin:[''],


        license_plate:[''],


      });


  }




  submit(){


    if(
      this.vehicleForm.invalid
    ){

      return;

    }



    this.service

      .createVehicle(

        this.vehicleForm.value

      )

      .subscribe({

        next:()=>{


          this.router.navigate([
            '/vehicles'
          ]);


        },


        error:err=>{

          console.error(err);

        }


      });


  }


}