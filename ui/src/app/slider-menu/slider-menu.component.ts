import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';

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
  constructor(private timetableHttpService: TimetableHttpService, private timetableService: TimetableService) {}
  ngOnInit(): void {
    this.timetableHttpService.getData(this.groupsUrl).subscribe(groups => (this.groups = groups));
    this.timetableHttpService.getData(this.subjectsUrl).subscribe(subjects => (this.subjects = subjects));
    this.timetableHttpService.getData(this.roomsUrl).subscribe(rooms => (this.rooms = rooms));
    this.timetableHttpService
      .getData(this.usersUrl, '?role=lecturer')
      .subscribe(lecturers => (this.lecturers = lecturers));
  }
  changeStartDate(value: any) {
    this.lesson.startAt = value;
  }
  updateLesson(subject: iSubject, room: Room, lecturer: User, group: Group, startAt: Date, lessonNum: string) {
    this.timetableService.changeProgressBarState(true);
    this.lesson.subject = subject;
    this.lesson.room = room;
    this.lesson.lecturer = lecturer;
    this.lesson.group = group;
    this.lesson.startAt = new Date(startAt);
    this.lesson.lessonNum = lessonNum;
    this.timetableHttpService.updateLesson(this.lesson).subscribe(() => {
      this.timetableService.changeProgressBarState(false);
      this.hide();
    });
  }
  deleteLesson() {
    this.timetableService.changeProgressBarState(true);
    this.deletedLesson.emit(this.lesson);
    this.timetableHttpService.deleteLesson(this.lesson.id).subscribe(() => {
      this.timetableService.changeProgressBarState(false);
      this.hide();
    });
  }
  addNewLesson(subject: iSubject, room: Room, lecturer: User, group: Group, startAt: Date, lessonNum: string) {
    this.timetableService.changeProgressBarState(true);
    this.timetableHttpService
      .addLesson({subject, room, startAt, lessonNum, lecturer, group} as Lesson)
      .subscribe(lesson => {
        this.timetableService.changeProgressBarState(false);
        this.addLesson.emit(lesson);
      });
    this.hide();
  }
  hide() {
    this.sliderAddNewState = false;
    this.hideComponent.emit(true);
  }
}
