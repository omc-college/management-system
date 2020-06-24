import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {AuthorizationService} from '../authorization.service';
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
  constructor(private authorisationService: AuthorizationService) {}

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
        email: value.email,
        password: value.password,
      };
      this.executeSignIn(signInValue);
    }
  }

  private executeSignIn(value: SignIn) {
    this.authorisationService.signIn(value).subscribe(request => console.log(request));
  }

  hasSignInError(controlName: string, errorName: string) {
    return this.signInForm.controls[controlName].hasError(errorName);
  }
}
