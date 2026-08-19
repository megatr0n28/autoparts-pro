import { Component, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';

import { AdminService } from '../core/services/admin.service';
import { UpdateVehicleRequest, Vehicle } from '../core/models/vehicle.model';

@Component({
  selector: 'app-admin-vehicle-edit',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, MatButtonModule, MatCardModule, MatInputModule],
  templateUrl: './admin-vehicle-edit.html',
  styleUrl: './admin-vehicle-edit.scss',
})
export class AdminVehicleEditComponent implements OnInit {
  form: FormGroup;
  vehicle: Vehicle | null = null;
  loading = signal(true);
  saving = signal(false);
  error = signal('');

  constructor(
    private readonly fb: FormBuilder,
    private readonly adminService: AdminService,
    private readonly route: ActivatedRoute,
    private readonly router: Router,
  ) {
    this.form = this.fb.group({
      year: ['', [Validators.required, Validators.min(1900)]], make: ['', Validators.required], model: ['', Validators.required],
      trim: [''], engine: [''], drivetrain: [''], transmission: [''], mileage: ['', Validators.min(0)], color: [''],
      vin: [''], license_plate: [''], state: [''],
    });
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.adminService.getOverview().subscribe({
      next: overview => {
        this.vehicle = overview.vehicles.find(item => item.id === id) ?? null;
        if (!this.vehicle) this.error.set('Vehicle not found.');
        else this.form.patchValue(this.vehicle);
        this.loading.set(false);
      },
      error: err => {
        this.error.set(err.status === 403 ? 'Admin access is required.' : 'Unable to load vehicle.');
        this.loading.set(false);
      },
    });
  }

  save(): void {
    if (!this.vehicle || this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    const value = this.form.value;
    const request: UpdateVehicleRequest = { ...value, year: Number(value.year), mileage: Number(value.mileage || 0) };
    this.saving.set(true);
    this.adminService.updateVehicle(this.vehicle.id, request).subscribe({
      next: () => this.router.navigate(['/admin']),
      error: err => { this.error.set(err.error?.error ?? 'Unable to update vehicle.'); this.saving.set(false); },
    });
  }

  cancel(): void { this.router.navigate(['/admin']); }
}
