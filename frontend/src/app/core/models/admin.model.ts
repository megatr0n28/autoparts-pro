import { Vehicle } from './vehicle.model';

export interface AdminUser {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  role: string;
  active: boolean;
  created_at: string;
}

export interface AdminCustomer {
  id: string;
  user_id: string;
  first_name: string;
  last_name: string;
  phone: string;
  address_line1: string;
  address_line2: string;
  city: string;
  state: string;
  postal_code: string;
  country: string;
}

export interface AdminInvoice {
  id: string;
  status: string;
}

export interface AdminOverview {
  users: AdminUser[];
  customers: AdminCustomer[];
  vehicles: Vehicle[];
  invoices: AdminInvoice[];
  invoice_management_enabled: boolean;
}
