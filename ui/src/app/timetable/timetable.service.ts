import {Injectable} from '@angular/core';
import {Observable, Subject} from 'rxjs';

import {Lesson} from '../models/Lesson';
import {GroupAsResource} from '../models/GroupAsResource';
import {UserAsResource} from '../models/UserAsResource';
import {SubjectInterface} from '../models/Subject';
import {Room} from '../models/Room';

@Injectable({
  providedIn: 'root',
})
export class TimetableService {
  private AddLessonComponentState$: Subject<boolean> = new Subject<boolean>();
  private progressBarState$: Subject<boolean> = new Subject<boolean>();
  private searchResult$: Subject<Lesson[]> = new Subject<Lesson[]>();
  private selectedDate$: Subject<Date> = new Subject<Date>();

  private lecturers$: Subject<UserAsResource[]> = new Subject<UserAsResource[]>();
  private subjects$: Subject<SubjectInterface[]> = new Subject<SubjectInterface[]>();
  private groups$: Subject<GroupAsResource[]> = new Subject<GroupAsResource[]>();
  private rooms$: Subject<Room[]> = new Subject<Room[]>();

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

  setSubjects(subjects: SubjectInterface[]): void {
    this.subjects$.next(subjects);
  }

  getSubjects(): Observable<SubjectInterface[]> {
    return this.subjects$;
  }

  setRooms(rooms: Room[]): void {
    this.rooms$.next(rooms);
  }

  getRooms(): Observable<Room[]> {
    return this.rooms$;
  }

  setLecturers(lecturer: UserAsResource[]): void {
    this.lecturers$.next(lecturer);
  }

  getLecturers(): Observable<UserAsResource[]> {
    return this.lecturers$;
  }

  setGroups(groups: GroupAsResource[]): void {
    this.groups$.next(groups);
  }

  getGroups(): Observable<GroupAsResource[]> {
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
}
