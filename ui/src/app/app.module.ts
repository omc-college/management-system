import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {AppRoutingModule} from './app-routing.module';

import {TimetableModule} from './timetable/timetable.module';
import {SignInModule} from './sign-in/sign-in.module';
import {LandingPageModule} from './landing-page/landing-page.module';
import {ErrorPageModule} from './error-page/error-page.module';
<<<<<<< HEAD

import {AdminModule} from './admin/admin.module';

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import {AppComponent} from './app.component';
=======

import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatNativeDateModule} from '@angular/material/core';
import {MatInputModule} from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';
import {MatSliderModule} from '@angular/material/slider';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatButtonModule} from '@angular/material/button';
import {MatBottomSheetModule} from '@angular/material/bottom-sheet';
import {MatIconModule} from '@angular/material/icon';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatCardModule} from '@angular/material/card';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatDialogModule} from '@angular/material/dialog';
import {MatExpansionModule} from '@angular/material/expansion';
>>>>>>> issue-96, ui structure of component rewrited, created sign-in/up, error and landing pages

import {AppComponent} from './app.component';

@NgModule({
  declarations: [AppComponent],
  imports: [
    SignInModule,
    TimetableModule,
    LandingPageModule,
    ErrorPageModule,
<<<<<<< HEAD
    AdminModule,
=======

>>>>>>> issue-96, ui structure of component rewrited, created sign-in/up, error and landing pages
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
