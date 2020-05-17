import {Component, OnInit} from '@angular/core';
import {TimetableService} from '../timetable.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.sass'],
})
export class HeaderComponent implements OnInit {
  AddLessonBtnState: boolean = false;
  showProgressBar: boolean = false;
  constructor(private timetableService: TimetableService) {}

  ngOnInit(): void {
    this.timetableService.getAddLessonComponentState().subscribe(bool => (this.AddLessonBtnState = bool));
    this.timetableService.getProgressBarState().subscribe(bool => (this.showProgressBar = bool));
  }
  addLesson() {
    this.AddLessonBtnState = !this.AddLessonBtnState;
    this.timetableService.changeAddLessonComponentState(this.AddLessonBtnState);
  }
}
