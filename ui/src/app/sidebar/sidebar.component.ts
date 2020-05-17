import {Component, OnInit} from '@angular/core';
import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';
import {MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';

import {Group} from '../models/Group';
import {User} from '../models/User';
import {iSubject} from '../models/Subject';
import {Room} from '../models/Room';
import {Error} from '../models/Error';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.sass'],
})
export class SidebarComponent implements OnInit {
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';
  showProgressBar: boolean = false;
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(
    private timetableHttpService: TimetableHttpService,
    private timetableService: TimetableService,
    private _bottomSheet: MatBottomSheet,
  ) {}
  ngOnInit(): void {
    this.timetableHttpService.getData(this.groupsUrl).subscribe(groups => (this.groups = groups));
    this.timetableHttpService.getData(this.subjectsUrl).subscribe(subjects => (this.subjects = subjects));
    this.timetableHttpService.getData(this.roomsUrl).subscribe(rooms => (this.rooms = rooms));
    this.timetableHttpService
      .getData(this.usersUrl, '?role=lecturer')
      .subscribe(lecturers => (this.lecturers = lecturers));
  }

  showError() {
    this._bottomSheet.open(ErrorComponent);
  }
  setProgressBar() {
    this.showProgressBar = !this.showProgressBar;
    this.timetableService.changeProgressBarState(this.showProgressBar);
  }
}

@Component({
  selector: 'app-error',
  templateUrl: 'error.html',
})
export class ErrorComponent {
  error: Error = {
    id: '55',
    code: 505,
    message: 'Here will be errors or other messages',
  };
  constructor(private _bottomSheetRef: MatBottomSheetRef<ErrorComponent>) {}
  openLink(event: MouseEvent): void {
    this._bottomSheetRef.dismiss();
    event.preventDefault();
  }
}
