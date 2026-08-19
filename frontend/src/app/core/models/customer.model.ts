export interface UpdateCustomerRequest {
  first_name: string;
  last_name: string;
  phone?: string;
  address_line1?: string;
  address_line2?: string;
  city?: string;
  state?: string;
  postal_code?: string;
}
