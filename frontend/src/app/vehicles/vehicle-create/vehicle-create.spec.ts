import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VehicleCreateComponent } from './vehicle-create';

describe('VehicleCreate', () => {
  let component: VehicleCreateComponent;
  let fixture: ComponentFixture<VehicleCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [VehicleCreateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(VehicleCreateComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
