import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

import {SignIn} from '../models/signIn';
@Injectable({
  providedIn: 'root',
})
export class AuthorizationService {
  httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };
  constructor(private http: HttpClient) {}

  signIn(value: SignIn): Observable<SignIn> {
    return this.http.post<SignIn>('api/signIn', value, this.httpOptions);
  }
}
