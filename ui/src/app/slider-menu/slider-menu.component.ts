<<<<<<< HEAD
import {Component, OnInit, Input, Output, EventEmitter, Inject} from '@angular/core';
import * as moment from 'moment';

import {TimetableHttpService} from '../shared/timetable-http.service';
import {TimetableService} from '../timetable.service';
import {MatSidenav} from '@angular/material/sidenav';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';

import {Lesson} from '../models/Lesson';
=======
import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';
import {Lesson} from '../models/Lesson';
import {TimetableHttpService} from '../shared/timetable-http.service';
import * as moment from 'moment';

>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
  @Input() hide: MatSidenav;
  @Input() lesson: Lesson;
  @Input() sliderAddNewState: boolean;
  @Output() deletedLesson = new EventEmitter();
  @Output() addLesson = new EventEmitter();

=======
  @Input() lesson: Lesson;
  @Input() lessons: Lesson[];
  @Input() sliderAddNewState: boolean;
  @Output() hideComponent = new EventEmitter();
  @Output() newStartDate = new EventEmitter();
  @Output() deletedLesson = new EventEmitter();
  @Output() addLesson = new EventEmitter();
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  private groupsUrl = 'api/groups';
  private roomsUrl = 'api/rooms';
  private subjectsUrl = 'api/subjects';
  private usersUrl = 'api/users';
<<<<<<< HEAD
  private isDelete: boolean = false;

  isChange: boolean = false;
  addNewForm: FormGroup;

=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  groups: Group[] = [];
  lecturers: User[] = [];
  subjects: iSubject[] = [];
  rooms: Room[] = [];
<<<<<<< HEAD
  constructor(
    private timetableHttpService: TimetableHttpService,
    private timetableService: TimetableService,
    public dialog: MatDialog,
  ) {}
  ngOnInit(): void {
    this.timetableHttpService.getData(this.groupsUrl).subscribe(groups => (this.groups = groups));
    this.timetableHttpService.getData(this.subjectsUrl).subscribe(subjects => (this.subjects = subjects));
    this.timetableHttpService.getData(this.roomsUrl).subscribe(rooms => (this.rooms = rooms));
    this.timetableHttpService
      .getData(this.usersUrl, '?role=lecturer')
      .subscribe(lecturers => (this.lecturers = lecturers));
    this.addNewForm = new FormGroup({
      subjectFormControl: new FormControl('', [Validators.required]),
      roomFormControl: new FormControl('', [Validators.required]),
      groupFormControl: new FormControl('', [Validators.required]),
      lecturerFormControl: new FormControl('', [Validators.required]),
      startDateFormControl: new FormControl('', [Validators.required]),
      endDateFormControl: new FormControl('', [Validators.required]),
      startTimeFormControl: new FormControl('', [Validators.required]),
      endTimeFormControl: new FormControl('', [Validators.required]),
    });
  }
  dateFilter(d: Date | null): boolean {
    const day = (d || new Date()).getDay();
    return day !== 0;
  }

  showRequiredErrorMessage(message: string): string {
    return `${message} is required`;
  }

  hasError(controlName: string, errorName: string) {
    return this.addNewForm.controls[controlName].hasError(errorName);
  }

  addNewLesson(LessonFormValue) {
    if (this.addNewForm.valid) {
      this.executeAddNewLesson(LessonFormValue);
    }
  }
  private executeAddNewLesson(lessonFormValue) {
    this.timetableService.changeProgressBarState(true);
    let newLesson: Lesson = {
      subject: lessonFormValue.subjectFormControl,
      room: lessonFormValue.roomFormControl,
      lecturer: lessonFormValue.lecturerFormControl,
      group: lessonFormValue.groupFormControl,
      startAt: moment(
        `${lessonFormValue.startDateFormControl.getFullYear()}-${
          lessonFormValue.startDateFormControl.getMonth() + 1
        }-${lessonFormValue.startDateFormControl.getDate()}T${lessonFormValue.startTimeFormControl}`,
        'YYYY-MM-DDTHH:mm',
      ),
      endAt: moment(
        `${lessonFormValue.endDateFormControl.getFullYear()}-${
          lessonFormValue.endDateFormControl.getMonth() + 1
        }-${lessonFormValue.endDateFormControl.getDate()}T${lessonFormValue.endTimeFormControl}`,
        'YYYY-MM-DDTHH:mm',
      ),
    } as Lesson;
    this.timetableHttpService.addLesson(newLesson).subscribe(
      lesson => {
        this.timetableService.changeProgressBarState(false);
        let dialogRef = this.dialog.open(SuccessDialog, {
          height: '170px',
          width: '300px',
          disableClose: true,
          data: {},
        });
        dialogRef.afterClosed().subscribe();
        this.addLesson.emit(lesson);
        this.addNewForm.reset();
      },
      error => {
        //temporary as well
        this.timetableService.changeProgressBarState(false);
      },
    );
    this.hideComponent();
  }
  public onCancel() {
    this.hideComponent();
    this.addNewForm.reset();
  }
  updateLesson(
    subject: iSubject,
    room: Room,
    lecturer: User,
    group: Group,
    dateFrom: Date,
    dateTo: Date,
    timeFrom: string,
    timeTo: string,
  ) {
    this.timetableService.changeProgressBarState(true);
    if (subject && subject.id !== this.lesson.subject.id) {
      this.lesson.subject = subject;
      this.isChange = true;
    }
    if (room && room.id !== this.lesson.room.id) {
      this.lesson.room = room;
      this.isChange = true;
    }
    if (lecturer && lecturer.id !== this.lesson.lecturer.id) {
      this.lesson.lecturer = lecturer;
      this.isChange = true;
    }
    if (group && group.id !== this.lesson.group.id) {
      this.lesson.group = group;
      this.isChange = true;
    }
    if (dateFrom) {
      let date = moment(`${dateFrom.getFullYear()}-${dateFrom.getMonth() + 1}-${dateFrom.getDate()}`, 'YYYY-MM-DD');
      this.lesson.startAt.year(date.year()).month(date.month()).date(date.date());
      this.isChange = true;
    }
    if (dateTo) {
      let date = moment(`${dateTo.getFullYear()}-${dateTo.getMonth() + 1}-${dateTo.getDate()}`, 'YYYY-MM-DD');
      this.lesson.endAt.year(date.year()).month(date.month()).date(date.date());
      this.isChange = true;
    }
    if (timeFrom !== this.lesson.startAt.format('HH:mm')) {
      let time = moment(timeFrom, 'HH:mm');
      this.lesson.startAt.hour(time.hour()).minute(time.minute());
      this.isChange = true;
    }
    if (timeTo !== this.lesson.endAt.format('HH:mm')) {
      let time = moment(timeTo, 'HH:mm');
      this.lesson.endAt.hour(time.hour()).minute(time.minute());
      this.isChange = true;
    }
    if (this.isChange) {
      this.timetableHttpService.updateLesson(this.lesson).subscribe(() => {
        this.timetableService.changeProgressBarState(false);
        this.hideComponent();
        this.isChange = false;
      });
    } else {
      this.timetableService.changeProgressBarState(false);
    }
  }

  deleteLesson() {
    const dialogRef = this.dialog.open(DeleteDialog, {
      width: '250px',
      data: {isDelete: this.isDelete},
    });

    dialogRef.afterClosed().subscribe(result => {
      this.isDelete = result;
      if (this.isDelete) {
        this.executeDeleteLesson();
        this.isDelete = false;
      } else {
        console.log('canceled');
      }
    });
  }
  private executeDeleteLesson() {
    this.timetableService.changeProgressBarState(true);
    this.deletedLesson.emit(this.lesson);
    this.timetableHttpService.deleteLesson(this.lesson.id).subscribe(() => {
      this.timetableService.changeProgressBarState(false);
      this.hideComponent();
    });
  }
  hideComponent(): void {
    this.addNewForm.reset();
    this.sliderAddNewState = false;
    this.timetableService.changeAddLessonComponentState(this.sliderAddNewState);
    this.hide.close();
  }
}

@Component({
  selector: 'delete-dialog',
  templateUrl: 'delete-dialog.html',
})
export class DeleteDialog {
  constructor(public dialogRef: MatDialogRef<DeleteDialog>, @Inject(MAT_DIALOG_DATA) public data) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
@Component({
  selector: 'success-dialog',
  templateUrl: 'success-dialog.html',
})
export class SuccessDialog {
  constructor(public dialogRef: MatDialogRef<SuccessDialog>, @Inject(MAT_DIALOG_DATA) public data) {}

  onNoClick(): void {
    this.dialogRef.close();
=======
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
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  }
}
