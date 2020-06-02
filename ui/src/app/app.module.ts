import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {AppRoutingModule} from './app-routing.module';

import {TimetableModule} from './timetable/timetable.module';
import {SignInModule} from './sign-in/sign-in.module';
import {LandingPageModule} from './landing-page/landing-page.module';
import {ErrorPageModule} from './error-page/error-page.module';

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import {AdminModule} from './admin/admin.module';

import {AppComponent} from './app.component';

@NgModule({
  declarations: [AppComponent],
  imports: [
    SignInModule,
    TimetableModule,
    LandingPageModule,
    ErrorPageModule,

    AdminModule,

    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,

    BrowserAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
