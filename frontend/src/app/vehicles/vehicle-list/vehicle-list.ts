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
  MatTableDataSource,
} from '@angular/material/table';

import {
  MatCardModule,
} from '@angular/material/card';

import {
  VehicleService,
} from '../../core/services/vehicle.service';


import {
  Vehicle,
} from '../../core/models/vehicle.model';

import {
  ActivatedRoute,
} from '@angular/router';


@Component({

  selector: 'app-vehicle-list',

  standalone: true,

  imports: [
    CommonModule,
    MatTableModule,
    MatCardModule,
  ],

  templateUrl:
    './vehicle-list.html',

  styleUrl:
    './vehicle-list.scss',

})


export class VehicleListComponent
implements OnInit {


  dataSource =
    new MatTableDataSource<Vehicle>([]);

  success = '';


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
    private route:
      ActivatedRoute,

  ) {}



  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      this.success =
        params['message'] ?? '';
    });
    this.loadVehicles();
  }



  loadVehicles(): void {


    this.vehicleService

      .getVehicles()

      .subscribe({

        next: vehicles => {


          console.log(
            "Vehicle List Loaded",
            vehicles
          );


          this.dataSource.data =
            vehicles;


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