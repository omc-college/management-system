import {Component, OnInit, Inject} from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {TimetableHttpService} from '../shared/timetable-http.service';

import {Group} from '../models/Group';
import {User} from '../models/User';
import {iSubject} from '../models/Subject';
import {Room} from '../models/Room';
@Component({
  selector: 'app-search-result',
  templateUrl: './search-result.component.html',
  styleUrls: ['./search-result.component.sass'],
})
export class SearchResultComponent implements OnInit {
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(
    private timetableHttpService: TimetableHttpService,
    public dialogRef: MatDialogRef<SearchResultComponent>,
    @Inject(MAT_DIALOG_DATA) public data,
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
  }
  dateFilter = (d: Date | null): boolean => {
    const day = (d || new Date()).getDay();
    return day !== 0;
  };
  close(): void {
    this.dialogRef.close();
  }
}
