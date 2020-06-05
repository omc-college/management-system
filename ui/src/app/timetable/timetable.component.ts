<<<<<<< HEAD
import {Component, OnInit, ViewChild, ElementRef} from '@angular/core';
import * as moment from 'moment';
import {MatSidenav, MatDrawer} from '@angular/material/sidenav';

import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';

import {Lesson} from '../models/Lesson';
=======
import {Component, OnInit} from '@angular/core';
import * as moment from 'moment';
import {TimetableService} from '../timetable.service';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674

@Component({
  selector: 'app-timetable',
  templateUrl: './timetable.component.html',
  styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
<<<<<<< HEAD
  @ViewChild('drawer') public drawer: MatSidenav;

  private lessonsUrl = 'api/lessons';
  private timestampUrl = 'api/timestamp1';

  readonly DAYSTOSHOW: number = 6;
  sliderAddNewState: boolean = false;

  timestamp: string[] = [];
  cards: moment.Moment[];
  lessons: Lesson[] = [];

  selectedLesson: Lesson;
  selectedDate: moment.Moment = moment();
  constructor(private timetableService: TimetableService, private timetableHttpService: TimetableHttpService) {}

  ngOnInit(): void {
    this.getSelectedDate();
    this.selectLesson();
    this.getLessons();
    this.setDatesInCard();
    this.getTimestamp();
    this.getAddLessonComponentState();
  }
  addNewLesson(newLesson: Lesson): void {
    newLesson.startAt = moment(newLesson.startAt);
    newLesson.endAt = moment(newLesson.endAt);
    if (
      this.cards.find(
        d =>
          d.year() === newLesson.startAt.year() &&
          d.month() === newLesson.startAt.month() &&
          d.date() === newLesson.startAt.date(),
      )
    ) {
      this.lessons.push(newLesson);
    }
  }

  deleteLesson(lesson: Lesson): void {
    this.lessons = this.lessons.filter(l => l !== lesson);
  }

  setDatesInCard(): void {
    this.cards = [];
    for (let i = 1; i <= this.DAYSTOSHOW; i++) {
      this.cards.push(moment().week(this.selectedDate.week()).day(i));
    }
  }
  getLessons(): void {
    this.timetableHttpService.getData(this.lessonsUrl).subscribe(lessons => {
      lessons.forEach(el => {
        el.startAt = moment(el.startAt);
        el.endAt = moment(el.endAt);
      });
      this.lessons = lessons;
      this.selectLesson();
    });
  }

  getSelectedDate(): void {
    this.timetableService.getSelectedDate().subscribe(date => {
      this.selectedDate = moment(date);
      this.setDatesInCard();
      this.selectLesson();
    });
  }

  setSelectedDate(date: moment.Moment): void {
    this.timetableService.selectDate(new Date(date.format('YYYY-MM-DD')));
  }

  selectLesson(): void {
    this.selectedLesson = this.lessons.find(
      el =>
        el.startAt.year() === this.selectedDate.year() &&
        el.startAt.month() === this.selectedDate.month() &&
        el.startAt.date() === this.selectedDate.date(),
    );
  }
  getTimestamp(): void {
    this.timetableHttpService.getData(this.timestampUrl).subscribe(timesetamp => (this.timestamp = timesetamp));
  }

  getAddLessonComponentState(): void {
    this.timetableService.getAddLessonComponentState().subscribe(bool => {
      this.sliderAddNewState = bool;
      this.drawer.toggle(this.sliderAddNewState);
    });
  }
  showAddLessonComponent(date: moment.Moment): void {
    this.sliderAddNewState = true;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.selectedDate = date;
    this.setSelectedDate(date);
  }
  showLessonInfo(lesson: Lesson): void {
    this.setSelectedDate(lesson.startAt);
    this.sliderAddNewState = false;
    this.selectedLesson = lesson;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
  }

=======
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
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  trackByMethod(index: number, el: any): number {
    return el.id;
  }
}
