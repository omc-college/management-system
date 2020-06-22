import {Component, OnInit, Inject} from '@angular/core';
import {ResourcesService} from './resources.service';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {DeleteDialogComponent} from '../../timetable/slider-menu/slider-menu.component';
import {AddRoleDialogComponent} from '../roles/roles.component';

import {Resource} from '../../models/Resource';

@Component({
  selector: 'app-resources-management',
  templateUrl: './resources-management.component.html',
  styleUrls: ['./resources-management.component.sass'],
})
export class ResourcesManagementComponent implements OnInit {
  resources: Resource[];
  selectedResources: Resource[] = [];

  isDelete = false;
  constructor(private resourcesService: ResourcesService, public dialog: MatDialog) {}

  ngOnInit(): void {
    this.getResources();
  }

  getResources(): void {
    this.resourcesService.getResources().subscribe(res => (this.resources = res));
  }

  deleteResources(): void {
    if (this.selectedResources.length === 0) {
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
          this.executeDeleteResources();
          this.isDelete = false;
        }
      });
    }
  }

  private executeDeleteResources(): void {
    this.selectedResources.map(res => {
      this.resourcesService.deleteResources(res);
    });
    this.resources = this.resources.filter(el => !this.selectedResources.find(f => f.resourceId === el.resourceId));
    this.selectedResources = [];
  }

  updateResources() {
    if (this.selectedResources.length === 0) {
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
          this.executeUpdateResources();
          this.isDelete = false;
        }
      });
    }
  }

  private executeUpdateResources() {
    this.selectedResources.map(res => {
      this.resourcesService.updateResources(res);
    });
    this.selectedResources = [];
  }

  addNewResource() {
    const newResource = {
      resourceName: '',
      resourceDescription: '',
    } as Resource;
    const dialogRef = this.dialog.open(AddRoleDialogComponent, {
      width: '250px',
      data: {message: 'Name of resource:'},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        newResource.resourceName = result;
        const dialogRef2 = this.dialog.open(AddRoleDialogComponent, {
          width: '300px',
          data: {message: 'Description of resource:'},
        });

        dialogRef2.afterClosed().subscribe(description => {
          if (description) {
            newResource.resourceDescription = description;
          }
          this.executeAddNewResource(newResource);
        });
      }
    });
  }

  private executeAddNewResource(res) {
    this.resourcesService.addResources(res).subscribe(r => this.resources.push(r));
  }

  selectUnselectResource(resource) {
    if (this.selectedResources.find(f => f === resource)) {
      this.selectedResources = this.selectedResources.filter(f => f !== resource);
    } else {
      this.selectedResources.push(resource);
    }
  }

  trackById(index, item) {
    return item.id;
  }
}

@Component({
  selector: 'app-message-dialog',
  templateUrl: 'message.html',
})
export class MessageDialogComponent {
  constructor(public dialogRef: MatDialogRef<DeleteDialogComponent>, @Inject(MAT_DIALOG_DATA) public data) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
