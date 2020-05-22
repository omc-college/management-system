import {Component, OnInit, ViewChild, ElementRef} from '@angular/core';
import * as moment from 'moment';
import {MatSidenav, MatDrawer} from '@angular/material/sidenav';

import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';

import {Lesson} from '../models/Lesson';

@Component({
  selector: 'app-timetable',
  templateUrl: './timetable.component.html',
  styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
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

  trackByMethod(index: number, el: any): number {
    return el.id;
  }
}
