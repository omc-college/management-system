import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Lesson } from './lesson';
import { LESSONS } from './mock-lessons';

@Injectable({
   providedIn: 'root',
})
export class TimetableService {
   constructor() {}
   getLessons(): Observable<Lesson[]> {
      return of(LESSONS);
   }
}
