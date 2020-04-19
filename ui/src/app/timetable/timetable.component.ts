import {Component, OnInit, Input} from '@angular/core';
import * as moment from 'moment';
import {TimetableService} from '../timetable.service';
import {Lesson} from '../iLesson';

@Component({
  selector: 'app-timetable',
  templateUrl: './timetable.component.html',
  styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
  lessons: Lesson[];
  selectedDate: moment.Moment = moment();
  showFirstShift: boolean = true;
  timestamps: string[] = [];
  constructor(private timetableService: TimetableService) {}
  ngOnInit(): void {
    this.getLessons();
    this.getSelectedDate();
    this.getTimestamp();
  }
  getLessons(): void {
    this.timetableService.getLessons().subscribe(lessons => {
      this.lessons = lessons;
      this.lessons.forEach(lesson => {
        lesson.startAt = moment(lesson.startAt);
        lesson.endAt = moment(lesson.endAt);
      });
    });
  }
  getSelectedDate(): void {
    this.timetableService.getSelectedDate().subscribe(date => (this.selectedDate = date));
  }
  setSelectedDate(y: number, m: number, d: number): void {
    this.selectedDate.year(y);
    this.selectedDate.month(m);
    this.selectedDate.date(d);
    this.timetableService.selectDate(this.selectedDate);
  }
  getTimestamp(): void {
    if (this.showFirstShift) {
      this.timetableService.getTimestamp1().subscribe(timesetamp => (this.timestamps = timesetamp));
    } else {
      this.timetableService.getTimestamp2().subscribe(timesetamp => (this.timestamps = timesetamp));
    }
    this.showFirstShift = !this.showFirstShift;
  }
  trackByMethod(index: number, el: any): number {
    return el.id;
  }
}
