import {Component, OnInit, Inject, ViewChild} from '@angular/core';

import {GroupsManagementService} from './groups-management.service';
import {DeleteDialogComponent} from '../../timetable/slider-menu/slider-menu.component';

import {MatTableDataSource} from '@angular/material/table';
import {MatPaginator} from '@angular/material/paginator';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';

import {Group} from '../../models/Group';
import {UserAsResource} from 'src/app/models/UserAsResource';

@Component({
  selector: 'app-groups-management',
  templateUrl: './groups-management.component.html',
  styleUrls: ['./groups-management.component.sass'],
})
export class GroupsManagementComponent implements OnInit {
  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  groups: Group[];
  dataSource = [];

  isDelete = false;

  displayedColumns: string[] = ['Name', 'LastName', 'Surname', 'Details', 'RemoveStudent'];

  constructor(private groupsManagementService: GroupsManagementService, public dialog: MatDialog) {}

  ngOnInit(): void {
    this.getGroups();
  }

  getGroups(): void {
    this.groupsManagementService.getGroups().subscribe(gr => {
      this.groups = gr;
      this.groups.map((group, id) => {
        this.dataSource[id] = new MatTableDataSource(group.students);
        this.dataSource[id].paginator = this.paginator;
      });
    });
  }

  deleteGroup(group: Group, id: number) {
    const dialogRef = this.dialog.open(DeleteDialogComponent, {
      width: '250px',
      data: {isDelete: this.isDelete},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.executeDeleteGroup(group, id);
      }
    });
  }

  private executeDeleteGroup(group: Group, id: number): void {
    this.groups = this.groups.filter(r => r !== group);
    this.groupsManagementService.deleteGroups(group).subscribe(result => {
      this.dataSource.splice(id, 1);
    });
  }

  addStudentsToGroup(group: Group, id: number): void {
    const dialogRef = this.dialog.open(SelectUserDialogComponent, {
      width: '80%',
      data: {students: group.students},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (result.length) {
          group.students = group.students.concat(result);
          this.updateGroup(group, id);
        }
      }
    });
  }

  delStudentsFromGroup(group: Group, student: UserAsResource, id: number) {
    const dialogRef = this.dialog.open(DeleteDialogComponent, {
      width: '250px',
      data: {isDelete: this.isDelete},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        group.students = group.students.filter(s => s.id !== student.id);
        this.updateGroup(group, id);
      }
    });
  }

  changeGroupInfo(group: Group, id: number): void {
    const dialogRef = this.dialog.open(GroupUpdateDialogComponent, {
      width: '60%',
      data: {
        specialisation: group.specialisation,
        yearOfEducation: group.yearOfEducation,
        groupNum: group.groupNumber,
        curator: group.curator,
        isAddNewGroup: false,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        if (result[0]) {
          group.specialisation = result[0];
        }
        if (result[1]) {
          group.yearOfEducation = result[1];
        }
        if (result[2]) {
          group.groupNumber = result[2];
        }
        if (result[3]) {
          group.curator = result[3];
        }
        this.updateGroup(group, id);
      }
    });
  }

  private updateGroup(group: Group, id: number) {
    this.groupsManagementService.updateGroup(group).subscribe(result => {
      this.dataSource[id] = new MatTableDataSource(group.students);
      this.dataSource[id].paginator = this.paginator;
    });
  }

  addNewGroup(): void {
    const dialogRef = this.dialog.open(GroupUpdateDialogComponent, {
      width: '60%',
      data: {
        isAddNewGroup: true,
      },
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const newGroup = {
          specialisation: result[0],
          yearOfEducation: result[1],
          groupNumber: result[2],
          curator: result[3],
        } as Group;

        const dialogRef2 = this.dialog.open(SelectUserDialogComponent, {
          width: '80%',
          data: {query: '?role=student'},
        });

        dialogRef2.afterClosed().subscribe(result2 => {
          if (result2) {
            newGroup.students = result2;
          }
          this.createGroup(newGroup);
        });
      }
    });
  }

  private createGroup(group: Group): void {
    this.groupsManagementService.addGroups(group).subscribe(result => {
      this.groups.push(result);
      this.dataSource.push(new MatTableDataSource(result.students));
    });
  }

  applyFilter(event: Event, id: number) {
    if (!(event.target as HTMLInputElement).value) {
      const filterValue = (event.target as HTMLInputElement).value;
      this.dataSource[id].filter = filterValue.trim().toLowerCase();
    } else {
      const filterValue = (event.target as HTMLInputElement).value;
      this.dataSource[id].filter = filterValue.trim().toLowerCase();
    }
  }

  trackById(index, item) {
    return item.id;
  }
}

// ---------------------------------------select user-------------------------

@Component({
  selector: 'app-select-users-dialog',
  templateUrl: 'user-list.html',
})
export class SelectUserDialogComponent implements OnInit {
  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  users: UserAsResource[];
  selectedUsers: UserAsResource[] = [];
  dataSource;

  isChecked = false;
  isDelete = false;
  isDisabled = false;

  displayedColumns: string[] = ['Checkbox', 'Name', 'LastName', 'Surname', 'Details'];

  constructor(
    private groupsManagementService: GroupsManagementService,
    public dialogRef: MatDialogRef<DeleteDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data,
  ) {}
  ngOnInit(): void {
    this.getUsers();
  }

  getUsers(): void {
    this.groupsManagementService.getUsersAsResource('?role=student').subscribe(us => {
      this.users = us;
      if (this.data.students) {
        this.data.students.map(s => (this.users = this.users.filter(u => u.id !== s.id)));
      }
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

  selectUnselectUser(User: UserAsResource): void {
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

  onNoClick(): void {
    this.dialogRef.close();
  }
}

// ---------------------------------------update/create group--------------------------

@Component({
  selector: 'app-update-group-dialog',
  templateUrl: 'group-update.html',
})
export class GroupUpdateDialogComponent implements OnInit {
  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  lecturers: UserAsResource[];
  specialisations: string[];
  yearsOfEducation: number[] = [1, 2, 3, 4];
  groupNums: number[] = [1, 2, 3];

  constructor(
    private groupsManagementService: GroupsManagementService,
    public dialogRef: MatDialogRef<DeleteDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data,
  ) {}

  ngOnInit(): void {
    this.getLecturers();
    this.getSpecialisations();
    this.yearsOfEducation = this.yearsOfEducation.filter(y => y !== this.data.yearOfEducation);
    this.groupNums = this.groupNums.filter(y => y !== this.data.groupNum);
  }

  getLecturers(): void {
    this.groupsManagementService.getUsersAsResource('?role=lecturer').subscribe(us => {
      this.lecturers = us;
      if (this.data.curator) {
        this.lecturers = this.lecturers.filter(l => l.id !== this.data.curator.id);
      }
    });
  }

  getSpecialisations(): void {
    this.groupsManagementService.getSpecialisations().subscribe(spec => {
      this.specialisations = spec;
      this.specialisations = this.specialisations.filter(f => f !== this.data.specialisation);
    });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  trackById(index, item) {
    return item.id;
  }
}
