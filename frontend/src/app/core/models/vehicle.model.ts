export interface Vehicle {

  id: string;

  customer_id: string;

  year: number;

  make: string;

  model: string;

  vin?: string;

  license_plate?: string;

  created_at?: string;

}



export interface CreateVehicleRequest {

  year: number;

  make: string;

  model: string;

  vin?: string;

  license_plate?: string;

}