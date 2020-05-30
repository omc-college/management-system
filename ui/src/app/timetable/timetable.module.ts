import {NgModule} from '@angular/core';
import {AppRoutingModule} from '../app-routing.module';

import {InMemoryDataService} from './in-memory-data.service';
import {TimetableComponent} from './timetable.component';
import {CalendarComponent} from './calendar/calendar.component';
import {HeaderComponent} from './header/header.component';
import {SidebarComponent} from './sidebar/sidebar.component';
import {ScheduleComponent} from './schedule/schedule.component';
import {SliderMenuComponent} from './slider-menu/slider-menu.component';
import {DeleteDialog} from './slider-menu/slider-menu.component';
import {SuccessDialog} from './slider-menu/slider-menu.component';
import {SearchResultComponent} from './search-result/search-result.component';

import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
// adding in-memory web-api
// Remove it when a real server is ready to receive requests.
import {HttpClientInMemoryWebApiModule} from 'angular-in-memory-web-api';

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

@NgModule({
  declarations: [
    ScheduleComponent,
    CalendarComponent,
    HeaderComponent,
    SidebarComponent,
    TimetableComponent,
    SliderMenuComponent,
    DeleteDialog,
    SuccessDialog,
    SearchResultComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    // The HttpClientInMemoryWebApiModule module intercepts HTTP requests
    // and returns simulated server responses.
    // Remove it when a real server is ready to receive requests.
    HttpClientInMemoryWebApiModule.forRoot(InMemoryDataService, {dataEncapsulation: false}),
    BrowserAnimationsModule,

    MatDatepickerModule,
    MatFormFieldModule,
    MatInputModule,
    MatNativeDateModule,
    MatSelectModule,
    MatSliderModule,
    MatSlideToggleModule,
    MatButtonModule,
    MatBottomSheetModule,
    MatIconModule,
    MatProgressBarModule,
    MatCardModule,
    MatSidenavModule,
    MatDialogModule,
    MatExpansionModule,
  ],
  exports: [
    ScheduleComponent,
    CalendarComponent,
    HeaderComponent,
    SidebarComponent,
    TimetableComponent,
    SliderMenuComponent,
    DeleteDialog,
    SuccessDialog,
    SearchResultComponent,
  ],
  entryComponents: [DeleteDialog, SuccessDialog],
  providers: [],
})
export class TimetableModule {}
