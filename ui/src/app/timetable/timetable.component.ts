import {Component, OnInit, ViewChild} from '@angular/core';
import {TimetableService} from '../timetable.service';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';
const NUMBEROFCARDS = 48;
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
  lessons: Lesson[] = [];
  selectedLesson: Lesson;
  selectedDate: Date = new Date();
  nextWeek: Date = new Date();
  constructor(private timetableService: TimetableService, private timetableHttpService: TimetableHttpService) {}
  ngOnInit(): void {
    this.getSelectedDate();
    this.getLessons();
    this.getTimestamp();
    this.getAddLessonComponentState();
  }
  getLessons(): void {
    this.timetableHttpService.getData(this.lessonsUrl).subscribe(lessons => {
      lessons.forEach(el => (el.startAt = new Date(el.startAt)));
      this.lessons = lessons;
    });
  }
  getSelectedDate(): void {
    this.timetableService.getSelectedDate().subscribe(date => {
      this.selectedDate = date;
    });
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
    this.lessons.forEach(el => {
      if (el === lesson) {
        el = null;
      }
    });
  }
  addNewLesson(newLesson: Lesson): void {
    newLesson.startAt = new Date(newLesson.startAt);
    this.lessons[newLesson.startAt.getDay() * 8 - (8 - +newLesson.lessonNum) - 1] = newLesson;
  }
  getAddLessonComponentState(): void {
    this.timetableService.getAddLessonComponentState().subscribe(bool => {
      this.sliderAddNewState = bool;
      this.hideSlider = !this.sliderAddNewState;
    });
  }
  showAddLessonComponent(): void {
    this.sliderAddNewState = true;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.hideSlider = false;
  }
  hideSliderComponent(hide: boolean) {
    this.sliderAddNewState = false;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.hideSlider = hide;
  }
  showSliderComponent(lesson: Lesson): void {
    this.selectedDate = lesson.startAt;
    this.timetableService.selectDate(this.selectedDate);
    this.sliderAddNewState = false;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.selectedLesson = lesson;
    this.hideSlider = false;
  }
  trackByMethod(index: number, el: any): number {
    return el.id;
  }
}
