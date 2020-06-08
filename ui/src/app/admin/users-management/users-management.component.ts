import {Component, OnInit, ViewChild} from '@angular/core';
import {UsersService} from './users.service';
import {MatDialog} from '@angular/material/dialog';
import * as moment from 'moment';
import {FormControl, FormGroup, Validators} from '@angular/forms';

import {MatTableDataSource} from '@angular/material/table';
import {MatPaginator} from '@angular/material/paginator';

import {DeleteDialogComponent} from '../../timetable/slider-menu/slider-menu.component';
import {MessageDialogComponent} from '../resources-management/resources-management.component';

import {User} from '../../models/User';
import {UrlResolver} from '@angular/compiler';

@Component({
  selector: 'app-users-management',
  templateUrl: './users-management.component.html',
  styleUrls: ['./users-management.component.sass'],
})
export class UsersManagementComponent implements OnInit {
  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  addNewForm: FormGroup;

  users: User[];
  selectedUsers: User[] = [];
  dataSource;

  isChecked = false;
  isDelete = false;
  isDisabled = false;

  displayedColumns: string[] = [
    'Checkbox',
    'Name',
    'LastName',
    'Surname',
    'DateOfBirth',
    'Email',
    'Phone',
    'CreatedAt',
    'ModifiedAt',
    'IsVerified',
    'Details',
  ];

  constructor(private usersService: UsersService, public dialog: MatDialog) {}

  ngOnInit(): void {
    this.getUsers();
    this.addNewForm = new FormGroup({
      firstNameFormControl: new FormControl('', [Validators.required]),
      lastNameFormControl: new FormControl('', [Validators.required]),
      surnameFormControl: new FormControl('', [Validators.required]),
      dateOfBirthFormControl: new FormControl('', [Validators.required]),
      emailFormControl: new FormControl('', [Validators.required, Validators.email]),
      mobilePhoneFormControl: new FormControl('', [Validators.required]),
    });
  }

  getUsers(): void {
    this.usersService.getUsers().subscribe(us => {
      this.users = us;
      this.dataSource = new MatTableDataSource(this.users);
      this.dataSource.paginator = this.paginator;
    });
  }

  deleteUsers(): void {
    if (this.selectedUsers.length === 0) {
      const dialogRef = this.dialog.open(MessageDialogComponent, {
        width: '250px',
        data: {message: 'You must select at least 1 element!'},
      });

      dialogRef.afterClosed().subscribe();
    } else {
      const dialogRef = this.dialog.open(DeleteDialogComponent, {
        width: '250px',
        data: {isDelete: this.isDelete},
      });

      dialogRef.afterClosed().subscribe(result => {
        this.isDelete = result;
        if (this.isDelete) {
          this.executeDeleteUsers();
          this.isDelete = false;
        }
      });
    }
  }

  private executeDeleteUsers(): void {
    this.selectedUsers.map(res => {
      this.usersService.deleteUsers(res);
    });
    this.users = this.users.filter(el => !this.selectedUsers.find(f => f.userId === el.userId));
    this.selectedUsers = [];
    this.dataSource = new MatTableDataSource(this.users);
    this.dataSource.paginator = this.paginator;
    this.isChecked = false;
  }

  addNewUser(userFormValue): void {
    if (this.addNewForm.valid) {
      this.executeAddNewUser(userFormValue);
    }
  }

  private executeAddNewUser(user): void {
    const newUser = {
      firstName: user.firstNameFormControl,
      lastName: user.lastNameFormControl,
      surname: user.surnameFormControl,
      dateOFBirth: moment(
        `${user.dateOfBirthFormControl.getFullYear()}-${
          user.dateOfBirthFormControl.getMonth() + 1
        }-${user.dateOfBirthFormControl.getDate()}`,
        'YYYY-MM-DD',
      ),
      email: user.emailFormControl,
      mobilePhone: user.mobilePhoneFormControl,
      createdAt: moment(),
      modifiedAt: moment(),
      roles: [],
      verified: false,
      userPhoto: '',
    } as User;
    this.usersService.addUsers(newUser).subscribe(u => {
      this.users.unshift(u);
      console.log(u);
      this.dataSource = new MatTableDataSource(this.users);
      this.dataSource.paginator = this.paginator;
    });
  }

  selectUnselectAllUsers(): void {
    if (!this.isChecked) {
      this.selectedUsers = this.users;
      this.isChecked = true;
    } else {
      this.selectedUsers = [];
      this.isChecked = false;
    }
  }

  selectUnselectUser(User: User): void {
    if (this.selectedUsers.find(u => u === User)) {
      this.selectedUsers = this.selectedUsers.filter(u => u !== User);
    } else {
      this.selectedUsers.push(User);
    }
  }

  applyFilter(event: Event) {
    if (!(event.target as HTMLInputElement).value) {
      const filterValue = (event.target as HTMLInputElement).value;
      this.dataSource.filter = filterValue.trim().toLowerCase();
      this.isDisabled = false;
      this.isChecked = false;
      this.selectedUsers = [];
    } else {
      const filterValue = (event.target as HTMLInputElement).value;
      this.dataSource.filter = filterValue.trim().toLowerCase();
      this.isDisabled = true;
      this.isChecked = false;
      this.selectedUsers = [];
    }
  }

  showRequiredErrorMessage(message: string): string {
    return `${message} is required`;
  }

  hasError(controlName: string, errorName: string) {
    return this.addNewForm.controls[controlName].hasError(errorName);
  }

  onCancel() {
    this.addNewForm.reset();
  }
}
