import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable, of} from 'rxjs';
import {Lesson} from '../models/Lesson';

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
    console.log('added');
    return this.http.post<Lesson>('api/lessons', lesson, this.httpOptions);
  }

  updateLesson(lesson: Lesson): Observable<any> {
    console.log('updated');
    return this.http.put(`api/lessons/${lesson.id}`, lesson, this.httpOptions);
  }

  deleteLesson(id: string): Observable<any> {
    console.log('deleted');
    return this.http.delete(`api/lessons/${id}`, this.httpOptions);
  }
}
