export interface LoginRequest {

  email: string;

  password: string;

}


export interface RegisterRequest {

  first_name: string;

  last_name: string;

  email: string;

  password: string;

}


export interface TokenResponse {

  access_token: string;

  refresh_token: string;

}


export interface CurrentUser {

  user_id: string;

  role: string;

}