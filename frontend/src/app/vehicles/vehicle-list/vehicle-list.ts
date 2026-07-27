import {
  Component,
  OnInit,
} from '@angular/core';


import {
  CommonModule,
} from '@angular/common';


import {
  MatTableModule,
} from '@angular/material/table';


import {
  VehicleService,
} from '../../core/services/vehicle.service';


import {
  Vehicle,
} from '../../core/models/vehicle.model';



@Component({

  selector: 'app-vehicle-list',

  standalone: true,

  imports: [

    CommonModule,

    MatTableModule,

  ],

  templateUrl:
    './vehicle-list.html',

  styleUrl:
    './vehicle-list.scss',

})


export class VehicleListComponent
implements OnInit {


  vehicles: Vehicle[] = [];


  displayedColumns = [

    'year',

    'make',

    'model',

    'vin',

    'license_plate',

  ];



  constructor(

    private vehicleService:
      VehicleService,

  ) {}



  ngOnInit(): void {


    this.loadVehicles();


  }



  loadVehicles() {


    this.vehicleService
      .getVehicles()

      .subscribe({

        next: data => {

          this.vehicles = data;

        },


        error: err => {

          console.error(
            "Failed loading vehicles",
            err
          );

        }

      });


  }


}