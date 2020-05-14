import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';
import * as moment from 'moment';

import {Group} from '../models/Group';
import {User} from '../models/User';
import {iSubject} from '../models/Subject';
import {Room} from '../models/Room';

@Component({
  selector: 'app-slider-menu',
  templateUrl: './slider-menu.component.html',
  styleUrls: ['./slider-menu.component.sass'],
})
export class SliderMenuComponent implements OnInit {
  @Input() lesson: Lesson;
  @Input() lessons: Lesson[];
  @Input() sliderAddNewState: boolean;
  @Output() hideComponent = new EventEmitter();
  @Output() newStartDate = new EventEmitter();
  @Output() deletedLesson = new EventEmitter();
  @Output() addLesson = new EventEmitter();
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
  addSubject: iSubject;
  addRoom: Room;
  addFrom: string;
  addTo: string;
  addTimeFrom: string;
  addTimeTo: string;
  addLecturer: User;
  addGroup: Group;
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
  changeStartDate(value: any) {
    this.lesson.startAt.year(moment(value, 'YYYY-MM-DD').year());
    this.lesson.startAt.month(moment(value, 'YYYY-MM-DD').month());
    this.lesson.startAt.date(moment(value, 'YYYY-MM-DD').date());
    this.newStartDate.emit(this.lesson.startAt);
  }
  changeEndDate(event: any) {
    this.lesson.endAt.year(moment(event, 'YYYY-MM-DD').year());
    this.lesson.endAt.month(moment(event, 'YYYY-MM-DD').month());
    this.lesson.endAt.date(moment(event, 'YYYY-MM-DD').date());
  }
  changeStartTime(value: any) {
    this.lesson.startAt.hour(moment(value, 'hh:mm').hour());
    this.lesson.startAt.minute(moment(value, 'hh:mm').minute());
  }
  changeEndTime(value: any) {
    this.lesson.endAt.hour(moment(value, 'hh:mm').hour());
    this.lesson.endAt.minute(moment(value, 'hh:mm').minute());
  }
  saveLesson() {
    this.timetableHttpService.updateLesson(this.lesson).subscribe(() => this.hide());
  }
  deleteLesson() {
    this.deletedLesson.emit(this.lesson);
    this.timetableHttpService.deleteLesson(this.lesson.id).subscribe(() => this.hide());
  }
  addNewLesson(subject, room, from, timeFrom, to, timeTo, [lecturers], [groups]) {
    let startAt: any = `${from}T${timeFrom}`;
    let endAt: any = `${to}T${timeTo}`;
    this.timetableHttpService
      .addLesson({subject, room, startAt, endAt, lecturers, groups} as Lesson)
      .subscribe(lesson => this.addLesson.emit(lesson));
    this.hide();
  }
  hide() {
    this.sliderAddNewState = false;
    this.hideComponent.emit(true);
  }
}
