import { Component, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';

import { AdminService } from '../core/services/admin.service';
import { AdminCustomer } from '../core/models/admin.model';
import { UpdateCustomerRequest } from '../core/models/customer.model';

@Component({
  selector: 'app-admin-customer-edit',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, MatButtonModule, MatCardModule, MatInputModule],
  templateUrl: './admin-customer-edit.html',
  styleUrl: './admin-customer-edit.scss',
})
export class AdminCustomerEditComponent implements OnInit {
  form: FormGroup;
  customer: AdminCustomer | null = null;
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
      first_name: ['', Validators.required], last_name: ['', Validators.required], phone: [''],
      address_line1: [''], address_line2: [''], city: [''], state: [''], postal_code: [''],
    });
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.adminService.getOverview().subscribe({
      next: overview => {
        this.customer = overview.customers.find(item => item.id === id) ?? null;
        if (!this.customer) this.error.set('Customer not found.');
        else this.form.patchValue(this.customer);
        this.loading.set(false);
      },
      error: err => {
        this.error.set(err.status === 403 ? 'Admin access is required.' : 'Unable to load customer.');
        this.loading.set(false);
      },
    });
  }

  save(): void {
    if (!this.customer || this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }
    this.saving.set(true);
    const request = this.form.value as UpdateCustomerRequest;
    this.adminService.updateCustomer(this.customer.id, request).subscribe({
      next: () => this.router.navigate(['/admin']),
      error: err => { this.error.set(err.error?.error ?? 'Unable to update customer.'); this.saving.set(false); },
    });
  }

  cancel(): void { this.router.navigate(['/admin']); }
}
