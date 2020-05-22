import {Component, OnInit} from '@angular/core';
import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';
import {MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';
import {FormControl, FormGroup} from '@angular/forms';
import {MatDialog} from '@angular/material/dialog';
import {SearchResultComponent} from '../search-result/search-result.component';

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
  private lessonsUrl = 'api/lessons';
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';

  showProgressBar: boolean = false;
  findForm: FormGroup;

  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(
    private timetableHttpService: TimetableHttpService,
    private timetableService: TimetableService,
    private _bottomSheet: MatBottomSheet,
    public dialog: MatDialog,
  ) {}

  ngOnInit(): void {
    this.timetableHttpService.getData(this.groupsUrl).subscribe(groups => {
      this.groups = groups;
    });

    this.timetableHttpService.getData(this.subjectsUrl).subscribe(subjects => {
      this.subjects = subjects;
    });

    this.timetableHttpService.getData(this.roomsUrl).subscribe(rooms => {
      this.rooms = rooms;
    });

    this.timetableHttpService.getData(this.usersUrl, '?role=lecturer').subscribe(lecturers => {
      this.lecturers = lecturers;
    });

    this.findForm = new FormGroup({
      subjectFormControl: new FormControl(''),
      roomFormControl: new FormControl(''),
      groupFormControl: new FormControl(''),
      lecturerFormControl: new FormControl(''),
      startDateFormControl: new FormControl(''),
      endDateFormControl: new FormControl(''),
      startTimeFormControl: new FormControl(''),
      endTimeFormControl: new FormControl(''),
    });
  }

  clear(): void {
    this.findForm.reset();
  }

  find(filters) {
    if (
      filters.subjectFormControl ||
      filters.roomFormControl ||
      filters.groupFormControl ||
      filters.lecturerFormControl ||
      filters.startDateFormControl ||
      filters.endDateFormControl ||
      filters.startTimeFormControl ||
      filters.endTimeFormControl
    ) {
      this.executeFind(filters);
    }
  }

  private executeFind(filters) {
    let query: string = `?`;
    if (filters.subjectFormControl && query.length === 1) {
      query += `subjectid=${filters.subjectFormControl.id}`;
    } else if (filters.subjectFormControl) {
      query += `&subjectid=${filters.subjectFormControl.id}`;
    }
    if (filters.roomFormControl && query.length === 1) {
      query += `roomid=${filters.roomFormControl.id}`;
    } else if (filters.roomFormControl) {
      query += `&roomid=${filters.roomFormControl.id}`;
    }
    if (filters.groupFormControl && query.length === 1) {
      query += `groupid=${filters.groupFormControl.id}`;
    } else if (filters.groupFormControl) {
      query += `&groupid=${filters.groupFormControl.id}`;
    }
    if (filters.lecturerFormControl && query.length === 1) {
      query += `lecturerid=${filters.lecturerFormControl.id}`;
    } else if (filters.lecturerFormControl) {
      query += `&lecturerid=${filters.lecturerFormControl.id}`;
    }
    if (filters.startDateFormControl && query.length === 1) {
      query += `datefrom=${filters.startDateFormControl.getFullYear()}-${filters.startDateFormControl.getMonth()}-${filters.startDateFormControl.getDate()}`;
    } else if (filters.startDateFormControl) {
      query += `&datefrom=${filters.startDateFormControl.getFullYear()}-${filters.startDateFormControl.getMonth()}-${filters.startDateFormControl.getDate()}`;
    }
    if (filters.endDateFormControl && query.length === 1) {
      query += `dateto=${filters.endDateFormControl.getFullYear()}-${filters.endDateFormControl.getMonth()}-${filters.endDateFormControl.getDate()}`;
    } else if (filters.endDateFormControl) {
      query += `&dateto=${filters.endDateFormControl.getFullYear()}-${filters.endDateFormControl.getMonth()}-${filters.endDateFormControl.getDate()}`;
    }
    if (filters.startTimeFormControl && query.length === 1) {
      query += `timefrom=${filters.startTimeFormControl}`;
    } else if (filters.startTimeFormControl) {
      query += `&timefrom=${filters.startTimeFormControl}`;
    }
    if (filters.endTimeFormControl && query.length === 1) {
      query += `timeto=${filters.endTimeFormControl}`;
    } else if (filters.endTimeFormControl) {
      query += `&timeto=${filters.endTimeFormControl}`;
    }
    console.log(query);
    this.timetableHttpService.getData(this.lessonsUrl, '?id=1').subscribe(lessons => {
      console.log(lessons);
      this.timetableService.setSearchResult(lessons);
      let dialogRef = this.dialog.open(SearchResultComponent, {
        height: '660px',
        width: '1240px',
        disableClose: true,
        data: {
          result: lessons,
        },
      });
      dialogRef.afterClosed().subscribe();
    });
  }

  dateFilter = (d: Date | null): boolean => {
    const day = (d || new Date()).getDay();
    return day !== 0;
  };
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
