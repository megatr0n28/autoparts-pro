export interface Vehicle {

  id: string;

  customer_id: string;

  vin: string;

  year: number;

  make: string;

  model: string;

  trim?: string;

  engine?: string;

  drivetrain?: string;

  transmission?: string;

  mileage?: number;

  color?: string;

  license_plate?: string;

  state?: string;

  is_primary?: boolean;

}



export interface CreateVehicleRequest {

  year: number;

  make: string;

  model: string;

  vin?: string;

  license_plate?: string;

}

export interface UpdateVehicleRequest {
  year: number;
  make: string;
  model: string;
  vin?: string;
  trim?: string;
  engine?: string;
  drivetrain?: string;
  transmission?: string;
  mileage?: number;
  color?: string;
  license_plate?: string;
  state?: string;
}