import { Component, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from '@angular/material/tabs';
import { RouterLink } from '@angular/router';

import { AdminOverview, AdminUser } from '../core/models/admin.model';
import { AdminService } from '../core/services/admin.service';

@Component({
  selector: 'app-admin',
  standalone: true,
  imports: [CommonModule, MatButtonModule, MatCardModule, MatTabsModule, RouterLink],
  templateUrl: './admin.html',
  styleUrl: './admin.scss',
})
export class AdminComponent implements OnInit {
  overview = signal<AdminOverview | null>(null);
  loading = signal(true);
  error = signal('');
  updatingUserId = signal('');

  constructor(private readonly adminService: AdminService) {}

  ngOnInit(): void {
    this.loadOverview();
  }

  loadOverview(): void {
    this.loading.set(true);
    this.adminService.getOverview().subscribe({
      next: overview => {
        this.overview.set(overview);
        this.loading.set(false);
      },
      error: err => {
        this.error.set(err.status === 403 ? 'Admin access is required.' : 'Unable to load admin data.');
        this.loading.set(false);
      },
    });
  }

  updateRole(user: AdminUser, role: string): void {
    if (user.role === role) {
      return;
    }

    this.updatingUserId.set(user.id);
    this.adminService.updateUserRole(user.id, role).subscribe({
      next: () => {
        user.role = role;
        this.updatingUserId.set('');
      },
      error: () => {
        this.error.set('Unable to update user access.');
        this.updatingUserId.set('');
      },
    });
  }
}
