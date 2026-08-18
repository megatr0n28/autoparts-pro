import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartsSearchComponent } from './parts-search';

describe('PartsSearch', () => {
  let component: PartsSearchComponent;
  let fixture: ComponentFixture<PartsSearchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PartsSearchComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PartsSearchComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
