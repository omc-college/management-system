import {TestBed} from '@angular/core/testing';

import {TimetableHttpService} from './timetable-http.service';

describe('TimetableHttpServiceService', () => {
  let service: TimetableHttpService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TimetableHttpService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
