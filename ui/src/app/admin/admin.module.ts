import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AppRoutingModule} from '../app-routing.module';
import {HttpClientModule} from '@angular/common/http';
import {HttpClientInMemoryWebApiModule} from 'angular-in-memory-web-api';
import {FormsModule} from '@angular/forms';
import {InMemoryDataService} from '../timetable/in-memory-data.service';

import {AdminComponent} from './admin.component';
import {RolesComponent, AddRoleDialogComponent, SelectFeatureDialogComponent} from './roles/roles.component';
import {UsersManagementComponent} from './users-management/users-management.component';
import {GroupsManagementComponent} from './groups-management/groups-management.component';
import {
  ResourcesManagementComponent,
  MessageDialogComponent,
} from './resources-management/resources-management.component';

import {MatTabsModule} from '@angular/material/tabs';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatCardModule} from '@angular/material/card';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatSortModule} from '@angular/material/sort';
import {MatTableModule} from '@angular/material/table';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatDialogModule} from '@angular/material/dialog';
import {MatCheckboxModule} from '@angular/material/checkbox';

@NgModule({
  declarations: [
    AdminComponent,
    RolesComponent,
    SelectFeatureDialogComponent,
    AddRoleDialogComponent,
    UsersManagementComponent,
    GroupsManagementComponent,
    ResourcesManagementComponent,
    MessageDialogComponent,
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    // The HttpClientInMemoryWebApiModule module intercepts HTTP requests
    // and returns simulated server responses.
    // Remove it when a real server is ready to receive requests.
    HttpClientInMemoryWebApiModule.forRoot(InMemoryDataService, {dataEncapsulation: false}),

    MatTabsModule,
    MatIconModule,
    MatButtonModule,
    MatTooltipModule,
    MatCardModule,
    MatExpansionModule,
    MatFormFieldModule,
    MatInputModule,
    MatSortModule,
    MatTableModule,
    MatToolbarModule,
    MatDialogModule,
    MatCheckboxModule,
  ],
  exports: [
    AdminComponent,
    RolesComponent,
    SelectFeatureDialogComponent,
    AddRoleDialogComponent,
    UsersManagementComponent,
    GroupsManagementComponent,
    ResourcesManagementComponent,
    MessageDialogComponent,

    MatTabsModule,
    MatIconModule,
    MatButtonModule,
    MatTooltipModule,
    MatCardModule,
    MatExpansionModule,
    MatFormFieldModule,
    MatInputModule,
    MatSortModule,
    MatTableModule,
    MatToolbarModule,
    MatDialogModule,
    MatCheckboxModule,
  ],
  providers: [],
})
export class AdminModule {}
