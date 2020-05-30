import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Lesson} from '../../models/Lesson';

@Injectable({
  providedIn: 'root',
})
export class TimetableHttpService {
  httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };

  constructor(private http: HttpClient) {}

  getData(url: string, query: string = ''): Observable<any[]> {
    return this.http.get<any[]>(`${url}${query}`);
  }

  addLesson(lesson: Lesson): Observable<Lesson> {
    return this.http.post<Lesson>('api/lessons', lesson, this.httpOptions);
  }

  updateLesson(lesson: Lesson): Observable<any> {
    return this.http.put(`api/lessons/${lesson.id}`, lesson, this.httpOptions);
  }

  deleteLesson(id: string): Observable<any> {
    return this.http.delete(`api/lessons/${id}`, this.httpOptions);
  }

  search(url: string, filters): Observable<any[]> {
    let params = new HttpParams();
    if (filters.subjectFormControl) {
      params = params.append('subjectid', filters.subjectFormControl.id);
    }
    if (filters.roomFormControl) {
      params = params.append('roomid', filters.roomFormControl.id);
    }
    if (filters.groupFormControl) {
      params = params.append('groupid', filters.groupFormControl.id);
    }
    if (filters.lecturerFormControl) {
      params = params.append('lecturerid', filters.lecturerFormControl.id);
    }
    if (filters.startDateFormControl) {
      params = params.append(
        'datefrom',
        `${filters.startDateFormControl.getFullYear()}-${filters.startDateFormControl.getMonth()}-${filters.startDateFormControl.getDate()}`,
      );
    }
    if (filters.endDateFormControl) {
      params = params.append(
        'dateto',
        `${filters.endDateFormControl.getFullYear()}-${filters.endDateFormControl.getMonth()}-${filters.endDateFormControl.getDate()}`,
      );
    }
    if (filters.startTimeFormControl) {
      params = params.append('timefrom', filters.startTimeFormControl);
    }
    if (filters.endTimeFormControl) {
      params = params.append('timeto', filters.endTimeFormControl);
    }
    return this.http.get<any[]>(url, {params: params});
  }
}
