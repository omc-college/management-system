import { Component, OnInit, Injectable } from '@angular/core';

@Component({
   selector: 'app-calendar',
   templateUrl: './calendar.component.html',
   styleUrls: ['./calendar.component.sass'],
})
@Injectable({
   providedIn: 'root',
})
export class CalendarComponent implements OnInit {
   readonly months: string[] = [
      'January',
      'February',
      'March',
      'April',
      'May',
      'June',
      'July',
      'August',
      'September',
      'October',
      'November',
      'December',
   ];
   days: Date[] = []; // needed for rendering days in month
   beforeDays: Date[] = []; // needed for rendering days before chosen month
   afterDays: Date[] = []; // needed for rendering days after chosen month
   switcher = false; // switch between showing days or months
   SelectedDate: Date = new Date(); // date which is chosen at this moment

   constructor() {}

   ngOnInit(): void {
      this.renderDays();
   }

   renderDays(): void {
      // fills days, beforeDays, AfterDays
      const daysInMonth: number = new Date(
         this.SelectedDate.getFullYear(),
         this.SelectedDate.getMonth() + 1,
         0
      ).getDate();
      // number of day in week of first day of month
      const startDay: number = new Date(this.SelectedDate.getFullYear(), this.SelectedDate.getMonth(), 1).getDay();

      const before: number = new Date(this.SelectedDate.getFullYear(), this.SelectedDate.getMonth(), 0).getDate() + 1;
      for (let i = startDay; i >= 0; i--) {
         this.beforeDays[startDay - i] = new Date(
            this.SelectedDate.getFullYear(),
            this.SelectedDate.getMonth() - 1,
            before - i
         );
      }
      // if we change year and previous date has more days than new date we just cut unuseful days
      this.beforeDays.splice(startDay);

      for (let day = 1; day <= daysInMonth; day++) {
         if (day === this.SelectedDate.getDate()) {
            this.days[day - 1] = this.SelectedDate;
            continue;
         }
         this.days[day - 1] = new Date(this.SelectedDate.getFullYear(), this.SelectedDate.getMonth(), day);
      }
      this.days.splice(daysInMonth);

      const after: number = 42 - daysInMonth - startDay;
      for (let i = 1; i <= after; i++) {
         this.afterDays[i - 1] = new Date(this.SelectedDate.getFullYear(), this.SelectedDate.getMonth() + 1, i);
      }
      this.afterDays.splice(after);
   }

   changeDay(e: Date): void {
      this.SelectedDate = e;
   }

   changeDate(y: number, m: number): void {
      this.SelectedDate = new Date(y, m, this.SelectedDate.getDate());
      this.renderDays();
   }
}
