import {Injectable} from '@angular/core';
import {Observable, of, Subject, GroupedObservable} from 'rxjs';
import {catchError, map, tap} from 'rxjs/operators';
import * as moment from 'moment';
import {HttpClient, HttpHeaders} from '@angular/common/http';

import {Lesson} from './iLesson';
import {Group} from './iGroup';
import {User} from './iUser';
import {iSubject} from './iSubject';
import {Room} from './iRoom';

@Injectable({
  providedIn: 'root',
})
export class TimetableService {
  selectedDate$: Subject<moment.Moment> = new Subject<moment.Moment>();
  private lessonsUrl = 'api/LESSONS';
  private groupsUrl = 'api/GROUPS';
  private roomsUrl = 'api/ROOMS';
  private subjectsUrl = 'api/SUBJECTS';
  private usersUrl = 'api/USERS';
  private timestamp1Url = 'api/TIMESTAMP1';
  private timestamp2Url = 'api/TIMESTAMP2';
  constructor(private http: HttpClient) {}
  getLessons(): Observable<Lesson[]> {
    return this.http.get<Lesson[]>(this.lessonsUrl).pipe(catchError(this.handleError<Lesson[]>('getLessons', [])));
  }
  getSelectedDate(): Observable<moment.Moment> {
    return this.selectedDate$;
  }
  selectDate(date: moment.Moment): void {
    this.selectedDate$.next(date);
  }
  getTimestamp1(): Observable<string[]> {
    return this.http.get<string[]>(this.timestamp1Url).pipe(catchError(this.handleError<string[]>('getLessons', [])));
  }
  getTimestamp2(): Observable<string[]> {
    return this.http.get<string[]>(this.timestamp2Url).pipe(catchError(this.handleError<string[]>('getLessons', [])));
  }
  getGroups(): Observable<Group[]> {
    return this.http.get<Group[]>(this.groupsUrl).pipe(catchError(this.handleError<Group[]>('getLessons', [])));
  }
  getLecturers(): Observable<User[]> {
    return this.http.get<User[]>(this.usersUrl).pipe(catchError(this.handleError<User[]>('getLessons', [])));
  }
  getSubjects(): Observable<iSubject[]> {
    return this.http.get<iSubject[]>(this.subjectsUrl).pipe(catchError(this.handleError<iSubject[]>('getLessons', [])));
  }
  getRooms(): Observable<Room[]> {
    return this.http.get<Room[]>(this.roomsUrl).pipe(catchError(this.handleError<Room[]>('getLessons', [])));
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      // TODO: send the error to remote logging infrastructure
      // log to console instead
      console.error(error, `${operation} failed: ${error.message}`);
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
