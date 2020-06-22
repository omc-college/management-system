import { TestBed } from '@angular/core/testing';

import { GroupsManagementService } from './groups-management.service';

describe('GroupsManagementService', () => {
  let service: GroupsManagementService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GroupsManagementService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
