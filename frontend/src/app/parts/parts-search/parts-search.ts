import {
  ChangeDetectorRef,
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

  templateUrl: './parts-search.html',
  styleUrl: './parts-search.scss',
})
export class PartsSearchComponent implements OnInit {

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
    private vehicleService: VehicleService,
    private partService: PartService,
    private cdr: ChangeDetectorRef,
  ) {}


  ngOnInit(): void {
    this.loadVehicles();
  }


  loadVehicles(): void {

    this.error = '';

    this.vehicleService
      .getVehicles()
      .subscribe({

        next: data => {
          this.vehicles = data ?? [];
        },

        error: err => {

          console.error(
            'Vehicle load error',
            err,
          );

          this.vehicles = [];

          this.error =
            'Unable to load vehicles. Please log in again.';

          this.cdr.detectChanges();
        },

      });
  }


  search(): void {

    this.error = '';

    this.results = [];


    if (!this.selectedVehicle) {

      this.error =
        'Please select a vehicle.';

      return;
    }


    const query =
      this.searchQuery.trim();


    if (!query) {

      this.error =
        'Please enter a part to search for.';

      return;
    }


    this.loading = true;


    this.partService
      .searchParts(
        this.selectedVehicle,
        query,
      )
      .subscribe({

        next: data => {

          this.results =
            data ?? [];

          this.loading = false;

          this.cdr.detectChanges();
        },


        error: err => {

          console.error(
            'Part search error',
            err,
          );

          this.results = [];

          this.error =
            err.error?.error ??
            'Search failed. Please try again.';

          this.loading = false;

          this.cdr.detectChanges();
        },

      });
  }
}