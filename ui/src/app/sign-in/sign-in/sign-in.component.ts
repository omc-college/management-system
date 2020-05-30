import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {AutorisationService} from '../autorisation.service';
import {SignIn} from '../../models/signIn';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.sass'],
})
export class SignInComponent implements OnInit {
  signInForm: FormGroup;
  signUpForm: FormGroup;
  showSignUp = false;
  hide = true;
  constructor(private autorisationService: AutorisationService) {}

  ngOnInit(): void {
    this.signInForm = new FormGroup({
      login: new FormControl('', [Validators.required, Validators.maxLength(60)]),
      password: new FormControl('', [Validators.required, Validators.maxLength(255)]),
    });
    this.signUpForm = new FormGroup({
      firstName: new FormControl('', [Validators.required, Validators.maxLength(255)]),
      lastname: new FormControl('', [Validators.required, Validators.maxLength(255)]),
      surname: new FormControl('', [Validators.required, Validators.maxLength(255)]),
      email: new FormControl('', [Validators.required, Validators.maxLength(60)]),
      password: new FormControl('', [Validators.required, Validators.maxLength(255)]),
    });
  }

  signIn(value) {
    if (this.signInForm.valid) {
      let signInValue: SignIn = {
        login: value.login,
        password: value.password,
      };
      this.executeSignIn(signInValue);
    }
  }

  private executeSignIn(value: SignIn) {
    this.autorisationService.signIn(value).subscribe(request => console.log(request));
  }

  signUp(value) {
    if (this.signUpForm.valid) {
      let signUpValue = {
        firstName: value.firstName,
        lastname: value.lastname,
        surname: value.surname,
        login: value.login,
        password: value.password,
      };
      this.executeSignUp(signUpValue);
    }
  }

  private executeSignUp(value) {
    this.autorisationService.signUp(value).subscribe(request => console.log(request));
  }
  hasSignInError(controlName: string, errorName: string) {
    return this.signInForm.controls[controlName].hasError(errorName);
  }
  hasSignUpError(controlName: string, errorName: string) {
    return this.signUpForm.controls[controlName].hasError(errorName);
  }
}
