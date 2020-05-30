import {NgModule} from '@angular/core';
import {Routes, RouterModule, CanActivate} from '@angular/router';

import {TimetableComponent} from './timetable/timetable.component';
import {SignInComponent} from './sign-in/sign-in/sign-in.component';
import {LandingPageComponent} from './landing-page/landing-page.component';
import {ErrorPageComponent} from './error-page/error-page.component';

class FutureGuard implements CanActivate {
  canActivate() {
    return true;
  }
}
const routes: Routes = [
  {path: '', redirectTo: '/landing', pathMatch: 'full'},
  {path: 'landing', component: LandingPageComponent, pathMatch: 'full'},
  {path: 'sign-in', component: SignInComponent, pathMatch: 'full'},
  {path: 'timetable', canActivate: [FutureGuard], component: TimetableComponent, pathMatch: 'full'},
  {path: '**', component: ErrorPageComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: [FutureGuard],
})
export class AppRoutingModule {}
