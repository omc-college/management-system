import {BrowserModule} from "@angular/platform-browser";
import {NgModule} from "@angular/core";

import {AppRoutingModule} from "./app-routing.module";
import {AppComponent} from "./app.component";
import {CalendarComponent} from "./calendar/calendar.component";
import {HeaderComponent} from "./header/header.component";
import {SidebarComponent} from "./sidebar/sidebar.component";
import {TimetableComponent} from "./timetable/timetable.component";

@NgModule({
   declarations: [AppComponent, CalendarComponent, HeaderComponent, SidebarComponent, TimetableComponent],
   imports: [BrowserModule, AppRoutingModule],
   providers: [],
   bootstrap: [AppComponent],
})
export class AppModule {}
