import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AppRoutingModule} from '../app-routing.module';

import {ErrorPageComponent} from './error-page.component';

import {MatCardModule} from '@angular/material/card';

@NgModule({
  declarations: [ErrorPageComponent],
  imports: [CommonModule, MatCardModule, AppRoutingModule],
  exports: [ErrorPageComponent],
})
export class ErrorPageModule {}
