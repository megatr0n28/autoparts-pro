import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../../environments/environment';
import { AdminOverview } from '../models/admin.model';
import { UpdateVehicleRequest } from '../models/vehicle.model';
import { UpdateCustomerRequest } from '../models/customer.model';

@Injectable({
  providedIn: 'root',
})
export class AdminService {
  private readonly api = environment.apiUrl;

  constructor(private readonly http: HttpClient) {}

  getOverview(): Observable<AdminOverview> {
    return this.http.get<AdminOverview>(`${this.api}/admin/overview`);
  }

  updateUserRole(id: string, role: string): Observable<void> {
    return this.http.put<void>(`${this.api}/admin/users/${id}/role`, { role });
  }

  updateVehicle(id: string, vehicle: UpdateVehicleRequest): Observable<void> {
    return this.http.put<void>(`${this.api}/admin/vehicles/${id}`, vehicle);
  }

  deleteVehicle(id: string): Observable<void> {
    return this.http.delete<void>(`${this.api}/admin/vehicles/${id}`);
  }

  updateCustomer(id: string, customer: UpdateCustomerRequest): Observable<void> {
    return this.http.put<void>(`${this.api}/admin/customers/${id}`, customer);
  }
}
