import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupsManagementComponent } from './groups-management.component';

describe('GroupsManagementComponent', () => {
  let component: GroupsManagementComponent;
  let fixture: ComponentFixture<GroupsManagementComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupsManagementComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupsManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
