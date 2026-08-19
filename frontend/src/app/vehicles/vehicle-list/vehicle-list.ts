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
  MatButtonModule,
} from '@angular/material/button';

import {
  VehicleService,
} from '../../core/services/vehicle.service';


import {
  Vehicle,
} from '../../core/models/vehicle.model';

import {
  ActivatedRoute,
  RouterLink,
} from '@angular/router';


@Component({

  selector: 'app-vehicle-list',

  standalone: true,

  imports: [
    CommonModule,
    MatTableModule,
    MatCardModule,
    MatButtonModule,
    RouterLink,
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

  selectedVehicle: Vehicle | null = null;

  private selectedVehicleId = '';

  success = '';


  displayedColumns = [

    'year',

    'make',

    'model',

    'vin',

    'license_plate',

    'actions',

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
      this.selectedVehicleId =
        params['vehicleId'] ?? '';
      this.selectVehicle();
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

          this.selectVehicle();


        },


        error: err => {


          console.error(
            "Failed loading vehicles",
            err
          );


        }

      });


  }

  private selectVehicle(): void {
    this.selectedVehicle =
      this.dataSource.data.find(
        vehicle => vehicle.id === this.selectedVehicleId,
      ) ?? null;
  }

  deleteVehicle(vehicle: Vehicle): void {
    const name = `${vehicle.year} ${vehicle.make} ${vehicle.model}`;

    if (!window.confirm(`Delete ${name}?`)) {
      return;
    }

    this.vehicleService.deleteVehicle(vehicle.id).subscribe({
      next: () => {
        this.dataSource.data =
          this.dataSource.data.filter(
            current => current.id !== vehicle.id,
          );

        if (this.selectedVehicle?.id === vehicle.id) {
          this.selectedVehicle = null;
        }

        this.success = 'Vehicle deleted successfully.';
      },
      error: err => {
        this.success =
          err.error?.error ?? 'Unable to delete vehicle.';
      },
    });
  }


}