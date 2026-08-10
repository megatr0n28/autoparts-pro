import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartsSearch } from './parts-search';

describe('PartsSearch', () => {
  let component: PartsSearch;
  let fixture: ComponentFixture<PartsSearch>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PartsSearch],
    }).compileComponents();

    fixture = TestBed.createComponent(PartsSearch);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
