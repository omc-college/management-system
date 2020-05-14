import {Component, OnInit} from '@angular/core';
import {Error} from '../models/Error';
import {TimetableHttpService} from '../shared/timetable-http.service';
import {Group} from '../models/Group';
import {User} from '../models/User';
import {iSubject} from '../models/Subject';
import {Room} from '../models/Room';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.sass'],
})
export class SidebarComponent implements OnInit {
  error: Error = {
    id: '55',
    code: 505,
    message: 'Here will be errors or other messages',
  };
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';
  showGroups: boolean = true;
  showLecturers: boolean = true;
  showSubjects: boolean = true;
  showRooms: boolean = true;
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(private timetableHttpService: TimetableHttpService) {}
  ngOnInit(): void {
    this.getGroups();
    this.getLecturers();
    this.getSubjects();
    this.getRooms();
  }
  getGroups(): void {
    this.timetableHttpService.getData(this.groupsUrl).subscribe(groups => (this.groups = groups));
  }
  getLecturers(): void {
    this.timetableHttpService
      .getData(this.usersUrl, '?role=lecturer')
      .subscribe(lecturers => (this.lecturers = lecturers));
  }
  getSubjects(): void {
    this.timetableHttpService.getData(this.subjectsUrl).subscribe(subjects => (this.subjects = subjects));
  }
  getRooms(): void {
    this.timetableHttpService.getData(this.roomsUrl).subscribe(rooms => (this.rooms = rooms));
  }
}
