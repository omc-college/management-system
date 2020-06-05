import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {HttpClientModule} from '@angular/common/http';
<<<<<<< HEAD
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
=======
import {FormsModule} from '@angular/forms';
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
// adding in-memory web-api
// Remove it when a real server is ready to receive requests.
import {HttpClientInMemoryWebApiModule} from 'angular-in-memory-web-api';
import {InMemoryDataService} from './in-memory-data.service';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {CalendarComponent} from './calendar/calendar.component';
import {HeaderComponent} from './header/header.component';
import {SidebarComponent} from './sidebar/sidebar.component';
import {TimetableComponent} from './timetable/timetable.component';
import {SliderMenuComponent} from './slider-menu/slider-menu.component';
<<<<<<< HEAD
import {DeleteDialog} from './slider-menu/slider-menu.component';
import {SuccessDialog} from './slider-menu/slider-menu.component';

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
import {SearchResultComponent} from './search-result/search-result.component';
import {MatExpansionModule} from '@angular/material/expansion';
=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674

@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent,
    HeaderComponent,
    SidebarComponent,
    TimetableComponent,
    SliderMenuComponent,
<<<<<<< HEAD
    DeleteDialog,
    SuccessDialog,
    SearchResultComponent,
=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
<<<<<<< HEAD
    ReactiveFormsModule,
=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
    // The HttpClientInMemoryWebApiModule module intercepts HTTP requests
    // and returns simulated server responses.
    // Remove it when a real server is ready to receive requests.
    HttpClientInMemoryWebApiModule.forRoot(InMemoryDataService, {dataEncapsulation: false}),
<<<<<<< HEAD
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
  entryComponents: [DeleteDialog, SuccessDialog],
=======
  ],
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
