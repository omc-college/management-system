import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {AuthorizationService} from '../authorization.service';
import {SignIn} from '../../models/signIn';
import {Title} from '@angular/platform-browser';
import {Meta} from '@angular/platform-browser';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.sass'],
})
export class SignInComponent implements OnInit {
  signInForm: FormGroup;

  hide = true;
  constructor(private authorisationService: AuthorizationService, private titleService: Title, private meta: Meta) {}

  ngOnInit(): void {
    this.titleService.setTitle('Authentication');
    this.meta.updateTag({name: 'description', content: 'Sign in/up to OMC learning management system'});

    this.signInForm = new FormGroup({
      login: new FormControl('', [Validators.required, Validators.minLength(8), Validators.maxLength(64)]),
      password: new FormControl('', [Validators.required, Validators.minLength(8), Validators.maxLength(64)]),
      isRememberMe: new FormControl(''),
    });
  }

  signIn(value) {
    if (this.signInForm.valid) {
      const signInValue: SignIn = {
        login: value.login,
        password: value.password,
        isRememberMe: value.isRememberMe,
      };
      this.executeSignIn(signInValue);
    }
  }

  private executeSignIn(value: SignIn) {
    this.authorisationService.signIn(value).subscribe();
  }

  hasSignInError(controlName: string, errorName: string) {
    return this.signInForm.controls[controlName].hasError(errorName);
  }
}
