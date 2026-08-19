import { Vehicle } from './vehicle.model';

export interface DashboardCustomer {
  id: string;
  first_name: string;
  last_name: string;
  phone: string;
  city: string;
  state: string;
}

export interface Dashboard {
  customer: DashboardCustomer;
  vehicles: Vehicle[];
  vehicle_count: number;
}