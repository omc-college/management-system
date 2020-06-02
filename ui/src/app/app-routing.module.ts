import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';

import {TimetableComponent} from './timetable/timetable.component';
import {SignInComponent} from './sign-in/sign-in/sign-in.component';
import {LandingPageComponent} from './landing-page/landing-page.component';
import {ErrorPageComponent} from './error-page/error-page.component';

import {AdminComponent} from './admin/admin.component';

const routes: Routes = [
  {path: '', redirectTo: '/landing', pathMatch: 'full'},
  {path: 'landing', component: LandingPageComponent, pathMatch: 'full'},
  {path: 'sign-in', component: SignInComponent, pathMatch: 'full'},
  {path: 'timetable', component: TimetableComponent, pathMatch: 'full'},
  {path: 'admin', component: AdminComponent, pathMatch: 'full'},
  {path: '**', component: ErrorPageComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
