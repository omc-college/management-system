import {Component, OnInit} from '@angular/core';
import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';
import {MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';
import {FormControl, FormGroup} from '@angular/forms';
import {MatDialog} from '@angular/material/dialog';
import {SearchResultComponent} from '../search-result/search-result.component';

import {GroupAsResource} from '../../models/GroupAsResource';
import {UserAsResource} from '../../models/UserAsResource';
import {iSubject} from '../../models/Subject';
import {Room} from '../../models/Room';
import {Error} from '../../models/Error';

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

  showProgressBar = false;
  findForm: FormGroup;

  groups: GroupAsResource[] = [];
  lecturers: UserAsResource[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(
    private timetableHttpService: TimetableHttpService,
    private timetableService: TimetableService,
    private bottomSheet: MatBottomSheet,
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
    )
      this.timetableHttpService.search(this.lessonsUrl, filters).subscribe(lessons => {
        this.timetableService.setSearchResult(lessons);
        const dialogRef = this.dialog.open(SearchResultComponent, {
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

  dateFilter(d: Date | null): boolean {
    const day = (d || new Date()).getDay();
    return day !== 0;
  }
  showError() {
    this.bottomSheet.open(ErrorComponent);
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

  constructor(private bottomSheetRef: MatBottomSheetRef<ErrorComponent>) {}

  openLink(event: MouseEvent): void {
    this.bottomSheetRef.dismiss();
    event.preventDefault();
  }
}
