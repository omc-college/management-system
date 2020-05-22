import {Component, OnInit, Input} from '@angular/core';
import {TimetableService} from '../timetable.service';
import {MatSidenav} from '@angular/material/sidenav';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.sass'],
})
export class HeaderComponent implements OnInit {
  @Input() inputSideNav: MatSidenav;
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
