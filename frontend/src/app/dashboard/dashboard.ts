import {
  Component,
  OnInit,
  signal,
} from '@angular/core';

import { DashboardService } from '../core/services/dashboard.service';
import { Dashboard } from '../core/models/dashboard.model';

import {
  MatCardModule,
} from '@angular/material/card';

import {
  MatButtonModule,
} from '@angular/material/button';

import { RouterLink } from '@angular/router';


@Component({

  selector: 'app-dashboard',

  standalone: true,

  imports: [

    MatCardModule,
    MatButtonModule,
    RouterLink,

  ],

  templateUrl: './dashboard.html',

  styleUrl: './dashboard.scss',

})

export class DashboardComponent implements OnInit {

  dashboard = signal<Dashboard | null>(null);

  loading = signal(true);

  error = signal('');

  constructor(
    private dashboardService: DashboardService,
  ) {}

  ngOnInit(): void {
    this.dashboardService.getDashboard().subscribe({
      next: dashboard => {
        this.dashboard.set(dashboard);
        this.loading.set(false);
      },
      error: () => {
        this.error.set('Unable to load dashboard data.');
        this.loading.set(false);
      },
    });
  }


}