import {Component, OnInit, ViewChild} from '@angular/core';
import {TimetableService} from '../timetable.service';

import {MatCalendar} from '@angular/material/datepicker';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.sass'],
})
export class CalendarComponent implements OnInit {
  selectedDate: Date = new Date(); // date which is chosen at this moment

  @ViewChild(MatCalendar) calendar: MatCalendar<Date>;

  constructor(private timetableService: TimetableService) {}

  ngOnInit(): void {
    this.getSelectedDate();
  }

  changeDate(date: Date): void {
    this.selectedDate = date;
    this.timetableService.selectDate(this.selectedDate);
  }

  getSelectedDate(): void {
    this.timetableService.getSelectedDate().subscribe(date => {
      this.selectedDate = date;
      this.calendar._goToDateInView(this.selectedDate, 'month');
    });
  }
}
