import { Component, OnInit, Input } from '@angular/core';
import { Lesson } from '../lesson';
import { LESSONS } from '../mock-lessons';
import { from } from 'rxjs';

@Component({
   selector: 'app-timetable',
   templateUrl: './timetable.component.html',
   styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
   @Input() selectedDate: Date;
   lessons: Lesson[] = LESSONS;
   selectedLesson: Date;
   constructor() {}
   ngOnInit(): void {
      this.selectLesson(this.selectedDate);
   }
   selectLesson(lid: Date): void {
      this.selectedLesson = lid;
      console.log(this.selectedLesson);
   }
}
