import {
  ComponentFixture,
  TestBed,
} from '@angular/core/testing';

import {
  HttpTestingController,
  provideHttpClientTesting,
} from '@angular/common/http/testing';

import {
  provideHttpClient,
} from '@angular/common/http';

import {
  PartsSearchComponent,
} from './parts-search';

import {
  VehicleService,
} from '../../core/services/vehicle.service';

import {
  Vehicle,
} from '../../core/models/vehicle.model';

import {
  PartService,
} from '../../core/services/part.service';

import {
  PartSearchResult,
} from '../../core/models/part.model';

describe('PartsSearch', () => {

  let component: PartsSearchComponent;

  let fixture: ComponentFixture<PartsSearchComponent>;

  let httpTestingController: HttpTestingController;


  const vehicle: Vehicle = {
    id: 'vehicle-1',
    customer_id: 'customer-1',
    vin: '1HGBH41JXMN109186',
    year: 2020,
    make: 'Toyota',
    model: 'Camry',
  };


  const searchResults: PartSearchResult[] = [
    {
      retailer: 'AutoZone',
      brand: 'FRAM',
      part_number: 'PH6607',
      name: 'FRAM Extra Guard Oil Filter',
      description: 'Engine oil filter',
      price: 9.99,
      currency: 'USD',
      in_stock: true,
      product_url: 'https://example.com/part',
      image_url: 'https://example.com/image.jpg',
    },
    {
      retailer: 'NAPA',
      brand: 'WIX',
      part_number: '57002',
      name: 'WIX Oil Filter',
      description: 'Premium oil filter',
      price: 12.49,
      currency: 'USD',
      in_stock: true,
      product_url: 'https://example.com/wix',
      image_url: 'https://example.com/wix.jpg',
    },
  ];


  beforeEach(async () => {

    await TestBed.configureTestingModule({

      imports: [
        PartsSearchComponent,
      ],

      providers: [
        VehicleService,
        PartService,

        provideHttpClient(),

        provideHttpClientTesting(),
      ],

    }).compileComponents();


    httpTestingController =
      TestBed.inject(
        HttpTestingController
      );


    fixture =
      TestBed.createComponent(
        PartsSearchComponent
      );


    component =
      fixture.componentInstance;


    fixture.detectChanges();


    const vehicleRequest =
      httpTestingController.expectOne(
        request =>
          request.url.includes(
            '/api/v1/vehicles'
          )
      );


    expect(
      vehicleRequest.request.method
    ).toBe('GET');


    vehicleRequest.flush([
      vehicle,
    ]);


    await fixture.whenStable();

  });


  afterEach(() => {

    httpTestingController.verify();

  });


  it('should create', () => {

    expect(component).toBeTruthy();

  });


  it('should load vehicles', () => {

    expect(
      component.vehicles
    ).toEqual([
      vehicle,
    ]);

  });


  it(
    'should require a vehicle before searching',
    () => {

      component.searchQuery =
        'oil filter';

      component.selectedVehicle =
        '';

      component.search();

      expect(
        component.error
      ).toBe(
        'Please select a vehicle.'
      );

    }
  );


  it(
    'should require a search query',
    () => {

      component.selectedVehicle =
        vehicle.id;

      component.searchQuery =
        '   ';

      component.search();

      expect(
        component.error
      ).toBe(
        'Please enter a part to search for.'
      );

    }
  );


  it(
    'should search for parts',
    () => {

      component.selectedVehicle =
        vehicle.id;

      component.searchQuery =
        'oil filter';


      component.search();


      const request =
        httpTestingController.expectOne(
          req =>
            req.url.includes(
              '/api/v1/parts/search'
            )
        );


      expect(
        request.request.method
      ).toBe('GET');


      expect(
        request.request.params.get(
          'vehicle_id'
        )
      ).toBe(
        vehicle.id
      );


      expect(
        request.request.params.get(
          'query'
        )
      ).toBe(
        'oil filter'
      );


      request.flush({
        results: searchResults,
      });


      expect(
        component.results
      ).toEqual(
        searchResults
      );


      expect(
        component.loading
      ).toBeFalsy();

    }
  );


  it(
    'should handle search errors',
    () => {

      component.selectedVehicle =
        vehicle.id;

      component.searchQuery =
        'oil filter';


      component.search();


      const request =
        httpTestingController.expectOne(
          req =>
            req.url.includes(
              '/api/v1/parts/search'
            )
        );


      request.flush(
        {
          error:
            'Part search failed',
        },
        {
          status: 500,
          statusText:
            'Internal Server Error',
        }
      );


      expect(
        component.results
      ).toEqual([]);


      expect(
        component.error
      ).toBe(
        'Part search failed'
      );


      expect(
        component.loading
      ).toBeFalsy();

    }
  );

  it('should sort search results by price', () => {

    component.selectedVehicle = vehicle.id;
    component.searchQuery = 'oil filter';

    component.search();

    const request =
      httpTestingController.expectOne(
        req =>
          req.url.includes(
            '/api/v1/parts/search'
          )
      );

    request.flush({
      results: [
        searchResults[1],
        searchResults[0],
      ],
    });

    expect(
      component.results[0]
    ).toEqual(
      searchResults[0]
    );

    expect(
      component.results[1]
    ).toEqual(
      searchResults[1]
    );

  });


  it('should identify the best price', () => {

    component.results =
      searchResults;

    expect(
      component.isBestPrice(
        searchResults[0]
      )
    ).toBe(true);

    expect(
      component.isBestPrice(
        searchResults[1]
      )
    ).toBe(false);

  });

});