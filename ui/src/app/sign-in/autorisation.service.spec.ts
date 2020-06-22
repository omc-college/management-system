import { TestBed } from '@angular/core/testing';

import { AutorisationService } from './autorisation.service';

describe('AutorisationService', () => {
  let service: AutorisationService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AutorisationService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
