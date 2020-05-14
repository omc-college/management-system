import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';
import * as moment from 'moment';
import {TimetableHttpService} from './shared/timetable-http.service';

@Injectable({
  providedIn: 'root',
})
export class TimetableService {
  selectedDate$: Subject<moment.Moment> = new Subject<moment.Moment>();
  AddLessonComponentState$: Subject<boolean> = new Subject<boolean>();

  constructor(private TimetableHttpService: TimetableHttpService) {}

  getSelectedDate(): Observable<moment.Moment> {
    return this.selectedDate$;
  }
  selectDate(date: moment.Moment): void {
    this.selectedDate$.next(date);
  }
  changeAddLessonComponentState(bool: boolean): void {
    this.AddLessonComponentState$.next(bool);
  }
  getAddLessonComponentState(): Observable<boolean> {
    return this.AddLessonComponentState$;
  }
}
