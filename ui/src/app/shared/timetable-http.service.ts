import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
<<<<<<< HEAD
import {Observable, of} from 'rxjs';
=======
import {Observable} from 'rxjs';
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
import {Lesson} from '../models/Lesson';

@Injectable({
  providedIn: 'root',
})
export class TimetableHttpService {
  httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };
<<<<<<< HEAD

  constructor(private http: HttpClient) {}

  getData(url: string, query: string = ''): Observable<any[]> {
    return this.http.get<any[]>(`${url}${query}`);
  }

=======
  constructor(private http: HttpClient) {}
  getData(url: string, query: string = ''): Observable<any[]> {
    return this.http.get<any[]>(`${url}${query}`);
  }
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  addLesson(lesson: Lesson): Observable<Lesson> {
    console.log('added');
    return this.http.post<Lesson>('api/lessons', lesson, this.httpOptions);
  }
<<<<<<< HEAD

=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  updateLesson(lesson: Lesson): Observable<any> {
    console.log('updated');
    return this.http.put(`api/lessons/${lesson.id}`, lesson, this.httpOptions);
  }
<<<<<<< HEAD

=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  deleteLesson(id: string): Observable<any> {
    console.log('deleted');
    return this.http.delete(`api/lessons/${id}`, this.httpOptions);
  }
}
