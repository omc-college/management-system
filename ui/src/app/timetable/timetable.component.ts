import { Component, OnInit, Injectable } from '@angular/core';
import { TimetableService } from '../timetable.service';
import { Lesson } from '../lesson';

@Component({
   selector: 'app-timetable',
   templateUrl: './timetable.component.html',
   styleUrls: ['./timetable.component.sass'],
})
@Injectable({
   providedIn: 'root',
})
export class TimetableComponent implements OnInit {
   selectedDate: Date;
   lessons: Lesson[];
   selectedLesson: Date;
   constructor(private timetableService: TimetableService) {}
   ngOnInit(): void {
      this.selectLesson(this.selectedDate);
      this.getLessons();
   }
   getLessons(): void {
      this.timetableService.getLessons().subscribe((lessons) => (this.lessons = lessons));
   }
   selectLesson(lid: Date): void {
      this.selectedLesson = lid;
      console.log(this.selectedLesson);
   }
}
