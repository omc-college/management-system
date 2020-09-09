import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {RegistrationService} from '../registration.service';
import {SignUp} from '../../models/signUp';

@Component({
   selector: 'app-sign-up',
   templateUrl: './sign-up.component.html',
   styleUrls: ['./sign-up.component.sass'],
})

export class SignUpComponent implements OnInit {

   signUpForm: FormGroup;
   hidePassword = true;

   constructor(private registrationService: RegistrationService) {}

   ngOnInit(): void {
      this.signUpForm = new FormGroup({
         first_name: new FormControl('', [Validators.required, Validators.maxLength(255)]),
         last_name: new FormControl('', [Validators.required, Validators.maxLength(255)]),
         email: new FormControl('', [Validators.required, Validators.maxLength(60)]),
         password: new FormControl('', [Validators.required, Validators.maxLength(255)]),
      });
   }

   signUp(value) {
      if (this.signUpForm.valid) {
         let signUpValue: SignUp = {
            first_name: value.first_name,
            last_name: value.last_name,
            email: value.email,
            password: value.password,
         };
         this.registrationService.signUp(signUpValue).subscribe(request => console.log(request),
               error => {console.log(error); });
      }
   }

   hasSignUpError(controlName: string, errorName: string) {
      return this.signUpForm.controls[controlName].hasError(errorName);
   }

}
