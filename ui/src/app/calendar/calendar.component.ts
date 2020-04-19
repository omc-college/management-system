import {Component, OnInit} from '@angular/core';
import * as moment from 'moment';
import {TimetableService} from '../timetable.service';

const CALENDAR_SIZE = 42;

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.sass'],
})
export class CalendarComponent implements OnInit {
  months: string[] = moment.months();
  showMonths: boolean = false; // switch between showing days or months
  selectedDate: moment.Moment = moment(); // date which is chosen at this moment
  daysForRender: moment.Moment[] = []; // needed for rendering days in month

  constructor(private timetableService: TimetableService) {}

  ngOnInit(): void {
    this.calculateDaysForRender();
    this.getSelectedDate();
  }

  calculateDaysForRender() {
    const firstDayOfSelectedMonth: number = moment()
      .year(this.selectedDate.year())
      .month(this.selectedDate.month())
      .date(0)
      .day();

    // in moment.js if you enter negative number in .date() it shows you previous month's date,
    // if your number will be > than number of days in selected month it shows you date from next month,
    // with help of i = -firstDayOfSelectedMonth it calculates days from previous month in first week of chosen month,
    // when i > days in chosen week it'll calculate next month's date
    for (let i = -firstDayOfSelectedMonth, k = 0; i < CALENDAR_SIZE - firstDayOfSelectedMonth; i++, k++) {
      this.daysForRender[k] = moment().year(this.selectedDate.year()).month(this.selectedDate.month()).date(i);
    }
  }
  changeDate(y: number, m: number, d: number): void {
    this.selectedDate.year(y);
    this.selectedDate.month(m);
    this.selectedDate.date(d);
    this.calculateDaysForRender();
  }
  getSelectedDate(): void {
    this.timetableService.getSelectedDate().subscribe(date => {
      this.selectedDate = date;
      this.calculateDaysForRender();
    });
  }
  setSelectedLesson(): void {}
}
