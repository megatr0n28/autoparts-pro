import {
  Injectable,
} from '@angular/core';


import {
  HttpClient,
} from '@angular/common/http';


import {
  Observable,
  tap,
} from 'rxjs';


import {
  environment,
} from '../../../environments/environment';


import {
  LoginRequest,
  RegisterRequest,
  TokenResponse,
  CurrentUser,
} from '../models/auth.model';



@Injectable({
  providedIn: 'root',
})
export class AuthService {


  private api =
    environment.apiUrl;



  private tokenKey =
    'access_token';



  constructor(
    private http: HttpClient,
  ) {}



  login(
    request: LoginRequest,
  ): Observable<TokenResponse> {


    return this.http.post<TokenResponse>(

      `${this.api}/auth/login`,

      request,

    )

    .pipe(

      tap(response => {


        localStorage.setItem(

          this.tokenKey,

          response.access_token

        );


      })

    );


  }



  register(
    request: RegisterRequest,
  ) {


    return this.http.post(

      `${this.api}/auth/register`,

      request,

    );


  }



  getCurrentUser() {


    return this.http.get<CurrentUser>(

      `${this.api}/users/me`

    );


  }



  logout() {

    localStorage.removeItem(
      this.tokenKey
    );

  }



  get token(): string | null {


    return localStorage.getItem(
      this.tokenKey
    );


  }



  isAuthenticated(): boolean {


    return !!this.token;


  }


}