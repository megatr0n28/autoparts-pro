import {
  Component,
  OnInit,
} from '@angular/core';

import {
  CommonModule,
} from '@angular/common';

import {
  FormsModule,
} from '@angular/forms';

import {
  MatInputModule,
} from '@angular/material/input';

import {
  MatButtonModule,
} from '@angular/material/button';

import {
  MatSelectModule,
} from '@angular/material/select';

import {
  MatTableModule,
} from '@angular/material/table';

import {
  VehicleService,
} from '../../core/services/vehicle.service';

import {
  PartService,
} from '../../core/services/part.service';

import {
  Vehicle,
} from '../../core/models/vehicle.model';

import {
  PartSearchResult,
} from '../../core/models/part.model';


@Component({

  selector: 'app-parts-search',

  standalone: true,

  imports: [

    CommonModule,

    FormsModule,

    MatInputModule,

    MatButtonModule,

    MatSelectModule,

    MatTableModule,

  ],

  templateUrl:
    './parts-search.html',

  styleUrl:
    './parts-search.scss',

})


export class PartsSearchComponent
implements OnInit {


  vehicles: Vehicle[] = [];


  selectedVehicle = '';

  searchQuery = '';


  results: PartSearchResult[] = [];


  loading = false;

  error = '';


  displayedColumns = [

    'name',

    'retailer',

    'brand',

    'part_number',

    'price',

    'availability',

  ];


  constructor(

    private vehicleService:
      VehicleService,

    private partService:
      PartService,

  ) {}


  ngOnInit(): void {

    console.log(
      'PartsSearchComponent INIT'
    );

    this.loadVehicles();

  }


  loadVehicles(): void {

    console.log(
      'LOADING VEHICLES'
    );


    this.vehicleService

      .getVehicles()

      .subscribe({

        next: data => {

          console.log(
            'VEHICLES LOADED',
            data
          );


          this.vehicles = data;

        },


        error: err => {

          console.error(
            'VEHICLE LOAD ERROR',
            err
          );


          this.error =
            'Unable to load vehicles. Please log in again.';

        },

      });

  }


  search(): void {

    console.log(
      'SEARCH BUTTON CLICKED'
    );


    this.error = '';

    this.results = [];


    if (!this.selectedVehicle) {

      this.error =
        'Please select a vehicle.';

      return;

    }


    if (!this.searchQuery.trim()) {

      this.error =
        'Please enter a part to search for.';

      return;

    }


    this.loading = true;


    const query =
      this.searchQuery.trim();


    console.log(
      'SEARCH START',
      this.selectedVehicle,
      query
    );


    this.partService

      .searchParts(

        this.selectedVehicle,

        query

      )

      .subscribe({

        next: data => {

          console.log(
            'SEARCH RESPONSE',
            data
          );


          this.results = data ?? [];


          console.log(
            'RESULT COUNT',
            this.results.length
          );


          console.log(
            'RESULTS AFTER ASSIGN',
            this.results
          );


          this.loading = false;


          console.log(
            'LOADING AFTER SUCCESS',
            this.loading
          );

        },


        error: err => {

          console.error(
            'SEARCH ERROR',
            err
          );


          this.results = [];


          this.error =
            err.error?.error ??
            'Search failed. Please try again.';


          this.loading = false;


          console.log(
            'LOADING AFTER ERROR',
            this.loading
          );

        },

      });

  }

}
