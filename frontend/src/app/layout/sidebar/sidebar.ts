import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';

interface NavigationItem {
  label: string;
  route: string;
  icon: string;
}

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive,
    MatIconModule,
  ],
  templateUrl: './sidebar.html',
  styleUrl: './sidebar.scss',
})
export class SidebarComponent {
  readonly navigationItems: NavigationItem[] = [
    {
      label: 'Dashboard',
      route: '/dashboard',
      icon: 'dashboard',
    },
    {
      label: 'Vehicles',
      route: '/vehicles',
      icon: 'directions_car',
    },
    {
      label: 'Parts Search',
      route: '/parts/search',
      icon: 'search',
    },
    {
      label: 'Customers',
      route: '/customers',
      icon: 'person',
    },
    {
      label: 'Invoices',
      route: '/invoices',
      icon: 'receipt_long',
    },
    {
      label: 'Admin',
      route: '/admin',
      icon: 'admin_panel_settings',
    },
  ];
}