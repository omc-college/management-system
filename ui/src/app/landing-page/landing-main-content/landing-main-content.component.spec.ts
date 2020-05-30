import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LandingMainContentComponent } from './landing-main-content.component';

describe('LandingMainContentComponent', () => {
  let component: LandingMainContentComponent;
  let fixture: ComponentFixture<LandingMainContentComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LandingMainContentComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LandingMainContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
