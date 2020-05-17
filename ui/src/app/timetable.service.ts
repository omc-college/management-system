import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';
import * as moment from 'moment';
import {TimetableHttpService} from './shared/timetable-http.service';

@Injectable({
  providedIn: 'root',
})
export class TimetableService {
  private selectedDate$: Subject<Date> = new Subject<Date>();
  private AddLessonComponentState$: Subject<boolean> = new Subject<boolean>();
  private progressBarState$: Subject<boolean> = new Subject<boolean>();

  constructor(private TimetableHttpService: TimetableHttpService) {}

  getSelectedDate(): Observable<Date> {
    return this.selectedDate$;
  }
  selectDate(date: Date): void {
    this.selectedDate$.next(date);
  }
  changeAddLessonComponentState(bool: boolean): void {
    this.AddLessonComponentState$.next(bool);
  }
  getAddLessonComponentState(): Observable<boolean> {
    return this.AddLessonComponentState$;
  }
  changeProgressBarState(bool: boolean): void {
    this.progressBarState$.next(bool);
  }
  getProgressBarState(): Observable<boolean> {
    return this.progressBarState$;
  }
}
