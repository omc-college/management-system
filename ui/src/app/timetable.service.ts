import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';
<<<<<<< HEAD
import {TimetableHttpService} from './shared/timetable-http.service';

import {Lesson} from './models/Lesson';
import {Group} from './models/Group';
import {User} from './models/User';
import {iSubject} from './models/Subject';
import {Room} from './models/Room';

=======
import * as moment from 'moment';
import {TimetableHttpService} from './shared/timetable-http.service';

>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
@Injectable({
  providedIn: 'root',
})
export class TimetableService {
<<<<<<< HEAD
  private AddLessonComponentState$: Subject<boolean> = new Subject<boolean>();
  private progressBarState$: Subject<boolean> = new Subject<boolean>();
  private searchResult$: Subject<Lesson[]> = new Subject<Lesson[]>();
  private selectedDate$: Subject<Date> = new Subject<Date>();

  private lecturers$: Subject<User[]> = new Subject<User[]>();
  private subjects$: Subject<iSubject[]> = new Subject<iSubject[]>();
  private groups$: Subject<Group[]> = new Subject<Group[]>();
  private rooms$: Subject<Room[]> = new Subject<Room[]>();

  constructor(private TimetableHttpService: TimetableHttpService) {}

  getSelectedDate(): Observable<Date> {
    return this.selectedDate$;
  }

  selectDate(date: Date): void {
    this.selectedDate$.next(date);
  }

  getSearchResult(): Observable<Lesson[]> {
    return this.searchResult$;
  }

  setSearchResult(result): void {
    this.searchResult$.next(result);
  }

  setSubjects(subjects: iSubject[]): void {
    this.subjects$.next(subjects);
  }

  getSubjects(): Observable<iSubject[]> {
    return this.subjects$;
  }

  setRooms(rooms: Room[]): void {
    this.rooms$.next(rooms);
  }

  getRooms(): Observable<Room[]> {
    return this.rooms$;
  }

  setLecturers(lecturer: User[]): void {
    this.lecturers$.next(lecturer);
  }

  getLecturers(): Observable<User[]> {
    return this.lecturers$;
  }

  setGroups(groups: Group[]): void {
    this.groups$.next(groups);
  }

  getGroups(): Observable<Group[]> {
    return this.groups$;
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
=======
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
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
}
