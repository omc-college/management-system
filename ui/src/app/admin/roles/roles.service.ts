import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

import {Role, FeatureEntry} from '../../models/role';

@Injectable({
  providedIn: 'root',
})
export class RolesService {
  private rolesUrl = 'api/roles';
  private featuresUrl = 'api/featureEntry';
  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };

  constructor(private http: HttpClient) {}

  getRoles(): Observable<Role[]> {
    console.log('get');
    return this.http.get<Role[]>(this.rolesUrl);
  }

  addRoles(role: Role): Observable<Role> {
    console.log('added');
    return this.http.post<Role>(this.rolesUrl, role, this.httpOptions);
  }

  deleteRoles(role: Role): Observable<Role> {
    console.log('deleted');
    const url = `api/roles/${role.id}`;
    return this.http.delete<Role>(url, this.httpOptions);
  }

  updateRole(role): Observable<Role> {
    console.log('updated');
    return this.http.put<Role>(this.rolesUrl, role, this.httpOptions);
  }

  getFeatures(): Observable<FeatureEntry[]> {
    console.log('getFeatures');
    return this.http.get<FeatureEntry[]>(this.featuresUrl);
  }
}
