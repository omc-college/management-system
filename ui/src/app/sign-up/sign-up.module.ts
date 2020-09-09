import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ReactiveFormsModule} from '@angular/forms';
import {AppRoutingModule} from "../app-routing.module";

import {SignUpComponent} from './sign-up/sign-up.component';

import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatSelectModule} from '@angular/material/select';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatStepperModule} from '@angular/material/stepper';

@NgModule({
   declarations: [SignUpComponent],
   imports: [
      CommonModule,
      ReactiveFormsModule,
      AppRoutingModule,

      MatCardModule,
      MatFormFieldModule,
      MatInputModule,
      MatButtonModule,
      MatSidenavModule,
      MatSelectModule,
      MatIconModule,
      MatToolbarModule,
      MatStepperModule,
   ],
   exports: [SignUpComponent],
})

export class SignUpModule {}
