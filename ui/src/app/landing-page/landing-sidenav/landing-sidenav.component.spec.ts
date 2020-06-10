import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LandingSidenavComponent } from './landing-sidenav.component';

describe('LandingSidenavComponent', () => {
  let component: LandingSidenavComponent;
  let fixture: ComponentFixture<LandingSidenavComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LandingSidenavComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LandingSidenavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
