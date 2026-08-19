import {
  Injectable,
} from '@angular/core';


import {
  HttpClient,
} from '@angular/common/http';


import {
  Observable,
} from 'rxjs';


import {
  environment,
} from '../../../environments/environment';


import {
  Vehicle,
  CreateVehicleRequest,
  UpdateVehicleRequest,
} from '../models/vehicle.model';



@Injectable({
  providedIn: 'root',
})
export class VehicleService {

  private api = environment.apiUrl;

  constructor(
    private http: HttpClient,
  ) {}

  getVehicles(): Observable<Vehicle[]> {

    return this.http.get<Vehicle[]>(

      `${this.api}/vehicles`

    );

  }

  createVehicle(
    vehicle: CreateVehicleRequest
  ) {

    return this.http.post(

      `${this.api}/vehicles`,

      vehicle

    );

  }

  deleteVehicle(id: string): Observable<void> {
    return this.http.delete<void>(
      `${this.api}/vehicles/${id}`,
    );
  }

  updateVehicle(
    id: string,
    vehicle: UpdateVehicleRequest,
  ): Observable<void> {
    return this.http.put<void>(
      `${this.api}/vehicles/${id}`,
      vehicle,
    );
  }

}