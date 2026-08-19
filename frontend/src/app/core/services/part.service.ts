import {
  Injectable,
} from '@angular/core';

import {
  HttpClient,
  HttpParams,
} from '@angular/common/http';

import {
  Observable,
  map,
} from 'rxjs';

import {
  environment,
} from '../../../environments/environment';

import {
  PartSearchResult,
} from '../models/part.model';


interface PartSearchResponse {

  results: PartSearchResult[];

}


@Injectable({
  providedIn: 'root',
})


export class PartService {


  private api =
    environment.apiUrl;


  constructor(
    private http: HttpClient,
  ) {}


  searchParts(
    vehicleId: string,
    query: string,
  ): Observable<PartSearchResult[]> {

    const params = new HttpParams()
      .set('vehicle_id', vehicleId)
      .set('query', query);

    return this.http
      .get<PartSearchResponse>(
        `${this.api}/parts/search`,
        {
          params,
        }
      )
      .pipe(
        map(response =>
          response.results ?? []
        )
      );
  }

}