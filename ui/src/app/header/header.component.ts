<<<<<<< HEAD
import {Component, OnInit, Input} from '@angular/core';
import {TimetableService} from '../timetable.service';
import {MatSidenav} from '@angular/material/sidenav';
=======
import {Component, OnInit} from '@angular/core';
import {TimetableService} from '../timetable.service';

>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.sass'],
})
export class HeaderComponent implements OnInit {
<<<<<<< HEAD
  @Input() inputSideNav: MatSidenav;
  AddLessonBtnState: boolean = false;
  showProgressBar: boolean = false;

=======
  AddLessonBtnState: boolean = false;
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  constructor(private timetableService: TimetableService) {}

  ngOnInit(): void {
    this.timetableService.getAddLessonComponentState().subscribe(bool => (this.AddLessonBtnState = bool));
<<<<<<< HEAD
    this.timetableService.getProgressBarState().subscribe(bool => (this.showProgressBar = bool));
  }

=======
  }
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  addLesson() {
    this.AddLessonBtnState = !this.AddLessonBtnState;
    this.timetableService.changeAddLessonComponentState(this.AddLessonBtnState);
  }
}
