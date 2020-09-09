import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders, HttpErrorResponse} from '@angular/common/http';
import {Observable, throwError} from "rxjs";
import {catchError} from "rxjs/operators";

import {SignUp} from '../models/signUp';
import {environment} from '../../environments/environment';

@Injectable({
   providedIn: 'root',
})

export class RegistrationService {

   httpOptions = {
      headers: new HttpHeaders({'Content-Type': 'application/json'}),
   };

   private handleError(error: HttpErrorResponse) {
      if (error.error instanceof ErrorEvent) {
         // A client-side or network error occurred. Handle it accordingly.
         console.error('An error occurred:', error.error.message);
      } else {
         // The backend returned an unsuccessful response code.
         // The response body may contain clues as to what went wrong,
         console.error(
            `Backend returned code ${error.status}, ` +
            `body was: ${error.error}`);
      }
      // return an observable with a user-facing error message
      return throwError(
         'Something bad happened; please try again later.');
   }

   constructor(private http: HttpClient) {}

   signUp(value: SignUp): Observable<SignUp> {
      let endpoint = `http://${environment.api}/sign-up`;
      return this.http.post<SignUp>(endpoint, value, this.httpOptions);
   }

}
