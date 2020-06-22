import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AppRoutingModule} from '../app-routing.module';

import {LandingPageComponent} from './landing-page.component';
import {LandingHeaderComponent} from './landing-header/landing-header.component';
import {LandingSidenavComponent} from './landing-sidenav/landing-sidenav.component';
import {LandingMainContentComponent} from './landing-main-content/landing-main-content.component';

import {MatButtonModule} from '@angular/material/button';

@NgModule({
  declarations: [LandingPageComponent, LandingHeaderComponent, LandingSidenavComponent, LandingMainContentComponent],
  imports: [CommonModule, MatButtonModule, AppRoutingModule],
  exports: [
    LandingPageComponent,
    LandingHeaderComponent,
    LandingSidenavComponent,
    LandingMainContentComponent,
    MatButtonModule,
  ],
})
export class LandingPageModule {}
