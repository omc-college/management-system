import {Component, OnInit} from '@angular/core';
import {Error} from '../iError';
import {TimetableService} from '../timetable.service';
import {Group} from '../iGroup';
import {User} from '../iUser';
import {iSubject} from '../iSubject';
import {Room} from '../iRoom';

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
  showGroups: boolean = true;
  showLecturers: boolean = true;
  showSubjects: boolean = true;
  showRooms: boolean = true;
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  constructor(private timetableService: TimetableService) {}
  ngOnInit(): void {
    this.getGroups();
    this.getLecturers();
    this.getSubjects();
    this.getRooms();
  }
  getGroups(): void {
    this.timetableService.getGroups().subscribe(groups => (this.groups = groups));
  }
  getLecturers(): void {
    this.timetableService.getLecturers().subscribe(lecturers => (this.lecturers = lecturers));
  }
  getSubjects(): void {
    this.timetableService.getSubjects().subscribe(subjects => (this.subjects = subjects));
  }
  getRooms(): void {
    this.timetableService.getRooms().subscribe(rooms => (this.rooms = rooms));
  }
}
