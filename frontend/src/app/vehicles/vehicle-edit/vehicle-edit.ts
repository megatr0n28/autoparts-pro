import { CommonModule } from '@angular/common';
import { Component, OnInit, signal } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';

import {
  UpdateVehicleRequest,
  Vehicle,
} from '../../core/models/vehicle.model';
import { VehicleService } from '../../core/services/vehicle.service';

@Component({
  selector: 'app-vehicle-edit',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
  ],
  templateUrl: './vehicle-edit.html',
  styleUrl: './vehicle-edit.scss',
})
export class VehicleEditComponent implements OnInit {
  vehicleForm: FormGroup;
  vehicle: Vehicle | null = null;
  loading = signal(true);
  submitting = false;
  error = signal('');

  constructor(
    private readonly fb: FormBuilder,
    private readonly vehicleService: VehicleService,
    private readonly route: ActivatedRoute,
    private readonly router: Router,
  ) {
    this.vehicleForm = this.fb.group({
      year: ['', [Validators.required, Validators.min(1900), Validators.max(new Date().getFullYear())]],
      make: ['', [Validators.required, Validators.minLength(2)]],
      model: ['', [Validators.required, Validators.minLength(1)]],
      trim: [''],
      engine: [''],
      drivetrain: [''],
      transmission: [''],
      mileage: ['', [Validators.min(0)]],
      color: [''],
      vin: [''],
      license_plate: [''],
      state: [''],
    });
  }

  ngOnInit(): void {
    const vehicleId = this.route.snapshot.paramMap.get('id');
    if (!vehicleId) {
      this.error.set('Vehicle not found.');
      this.loading.set(false);
      return;
    }

    this.vehicleService.getVehicles().subscribe({
      next: vehicles => {
        this.vehicle = vehicles.find(current => current.id === vehicleId) ?? null;
        if (!this.vehicle) {
          this.error.set('Vehicle not found.');
        } else {
          this.vehicleForm.patchValue(this.vehicle);
        }
        this.loading.set(false);
      },
      error: () => {
        this.error.set('Unable to load vehicle. Please sign in again.');
        this.loading.set(false);
      },
    });
  }

  submit(): void {
    if (!this.vehicle || this.vehicleForm.invalid) {
      this.vehicleForm.markAllAsTouched();
      return;
    }

    this.error.set('');
    this.submitting = true;
    const value = this.vehicleForm.value;
    const request: UpdateVehicleRequest = {
      ...value,
      year: Number(value.year),
      mileage: Number(value.mileage || 0),
    };

    this.vehicleService.updateVehicle(this.vehicle.id, request).subscribe({
      next: () => {
        this.router.navigate(['/vehicles'], {
          queryParams: {
            vehicleId: this.vehicle?.id,
            message: 'Vehicle updated successfully',
          },
        });
      },
      error: err => {
        this.error.set(err.error?.error ?? 'Unable to update vehicle.');
        this.submitting = false;
      },
    });
  }

  cancel(): void {
    this.router.navigate(['/vehicles'], {
      queryParams: this.vehicle ? { vehicleId: this.vehicle.id } : {},
    });
  }
}
