import {Component, OnInit} from '@angular/core';
import * as moment from 'moment';
import {TimetableService} from '../timetable.service';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';

@Component({
  selector: 'app-timetable',
  templateUrl: './timetable.component.html',
  styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
  private lessonsUrl = 'api/lessons';
  private timestamp1Url = 'api/timestamp1';
  private timestamp2Url = 'api/timestamp2';
  timestamps: string[] = [];
  showFirstShift: boolean = true;
  hideSlider: boolean = true;
  sliderAddNewState: boolean = false;
  lessons: Lesson[];
  selectedLesson: Lesson;
  selectedDate: moment.Moment = moment();
  constructor(private timetableService: TimetableService, private timetableHttpService: TimetableHttpService) {}
  ngOnInit(): void {
    this.getLessons();
    this.getSelectedDate();
    this.getTimestamp();
    this.getAddLessonComponentState();
  }
  getLessons(): void {
    this.timetableHttpService.getData(this.lessonsUrl).subscribe(lessons => {
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
      this.timetableHttpService.getData(this.timestamp1Url).subscribe(timesetamp => (this.timestamps = timesetamp));
    } else {
      this.timetableHttpService.getData(this.timestamp2Url).subscribe(timesetamp => (this.timestamps = timesetamp));
    }
    this.showFirstShift = !this.showFirstShift;
  }
  deleteLesson(lesson: Lesson): void {
    this.lessons = this.lessons.filter(l => l !== lesson);
  }
  addNewLesson(newLesson: Lesson): void {
    newLesson.startAt = moment(newLesson.startAt);
    newLesson.endAt = moment(newLesson.endAt);
    this.lessons.push(newLesson);
  }
  getAddLessonComponentState(): void {
    this.timetableService.getAddLessonComponentState().subscribe(bool => {
      this.sliderAddNewState = bool;
      this.hideSlider = !this.sliderAddNewState;
    });
  }
  hideSliderComponent(hide: boolean) {
    this.sliderAddNewState = false;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.hideSlider = hide;
  }
  showSliderComponent(lesson: Lesson): void {
    this.sliderAddNewState = false;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.selectedLesson = lesson;
    this.hideSlider = false;
  }
  trackByMethod(index: number, el: any): number {
    return el.id;
  }
}
