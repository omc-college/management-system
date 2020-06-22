import {Component, OnInit, Inject} from '@angular/core';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import {RolesService} from './roles.service';

import {Role, FeatureEntry} from '../../models/role';
import {DeleteDialogComponent} from '../../timetable/slider-menu/slider-menu.component';

@Component({
  selector: 'app-roles',
  templateUrl: './roles.component.html',
  styleUrls: ['./roles.component.sass'],
})
export class RolesComponent implements OnInit {
  roles: Role[];
  features: FeatureEntry[];
  private isDelete = false;

  constructor(private roleService: RolesService, public dialog: MatDialog) {}

  ngOnInit(): void {
    this.getRoles();
    this.getFeatures();
  }

  getFeatures(): void {
    this.roleService.getFeatures().subscribe(features => (this.features = features));
  }

  getRoles(): void {
    this.roleService.getRoles().subscribe(roles => (this.roles = roles));
  }

  deleteRole(role: Role) {
    const dialogRef = this.dialog.open(DeleteDialogComponent, {
      width: '250px',
      data: {isDelete: this.isDelete},
    });

    dialogRef.afterClosed().subscribe(result => {
      this.isDelete = result;
      if (this.isDelete) {
        this.executeDeleteRole(role);
        this.isDelete = false;
      }
    });
  }

  private executeDeleteRole(role: Role): void {
    this.roles = this.roles.filter(r => r !== role);
    this.roleService.deleteRoles(role).subscribe();
  }

  deleteFeatureInRole(role: Role, feature: FeatureEntry): void {
    const dialogRef = this.dialog.open(DeleteDialogComponent, {
      width: '250px',
      data: {isDelete: this.isDelete},
    });

    dialogRef.afterClosed().subscribe(result => {
      this.isDelete = result;
      if (this.isDelete) {
        this.executeDeleteFeatureInRole(role, feature);
        this.isDelete = false;
      }
    });
  }

  private executeDeleteFeatureInRole(role: Role, feature: FeatureEntry): void {
    const selectedRole = this.roles.find(r => r === role);
    selectedRole.entries = selectedRole.entries.filter(f => f !== feature);
    this.roleService.updateRole(selectedRole);
  }

  showAddFeatureDialog(role: Role): void {
    const filteredFeatures = this.features.filter(el => !role.entries.find(f => f.id === el.id));
    const dialogRef = this.dialog.open(SelectFeatureDialogComponent, {
      width: '800px',
      data: {features: filteredFeatures},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        role.entries = role.entries.concat(result);
        this.updateRole(role);
      }
    });
  }

  private updateRole(role: Role): void {
    this.roleService.updateRole(role).subscribe();
  }

  addRole(): void {
    const role: Role = {
      name: '',
      entries: [],
    } as Role;
    const roleRef = this.dialog.open(AddRoleDialogComponent, {
      width: '300px',
      data: {message: 'Name:'},
    });
    roleRef.afterClosed().subscribe(result => {
      if (result) {
        role.name = result;
        const featuresRef = this.dialog.open(SelectFeatureDialogComponent, {
          width: '800px',
          data: {features: this.features},
        });

        featuresRef.afterClosed().subscribe(res => {
          if (res) {
            role.entries = role.entries.concat(res);
          }
          this.executeAddRole(role);
        });
      }
    });
  }

  private executeAddRole(role): void {
    this.roleService.addRoles(role).subscribe(r => this.roles.push(r));
  }

  changeName(role): void {
    const roleRef = this.dialog.open(AddRoleDialogComponent, {
      width: '300px',
      data: {message: 'New name is:'},
    });
    roleRef.afterClosed().subscribe(result => {
      if (result) {
        role.name = result;
        this.updateRole(role);
      }
    });
  }
}

@Component({
  selector: 'app-select-feature-dialog',
  templateUrl: 'select-feature.html',
})
export class SelectFeatureDialogComponent {
  selectedFeatures: FeatureEntry[] = [];

  constructor(public dialogRef: MatDialogRef<DeleteDialogComponent>, @Inject(MAT_DIALOG_DATA) public data) {}

  selectUnselectFeature(feature: FeatureEntry): void {
    if (this.selectedFeatures.find(f => f === feature)) {
      this.selectedFeatures = this.selectedFeatures.filter(f => f !== feature);
    } else {
      this.selectedFeatures.push(feature);
    }
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}

@Component({
  selector: 'app-add-role-dialog',
  templateUrl: 'add-role.html',
})
export class AddRoleDialogComponent {
  name: string;

  constructor(public dialogRef: MatDialogRef<DeleteDialogComponent>, @Inject(MAT_DIALOG_DATA) public data) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
