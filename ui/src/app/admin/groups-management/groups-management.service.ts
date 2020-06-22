import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

import {Group} from '../../models/Group';
import {UserAsResource} from '../../models/UserAsResource';

@Injectable({
  providedIn: 'root',
})
export class GroupsManagementService {
  private GroupsUrl = 'api/fullGroups';
  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };

  constructor(private http: HttpClient) {}

  getGroups(): Observable<Group[]> {
    console.log('getGroups');
    return this.http.get<Group[]>(this.GroupsUrl);
  }

  addGroups(group: Group): Observable<Group> {
    console.log('added');
    return this.http.post<Group>(this.GroupsUrl, group, this.httpOptions);
  }

  updateGroup(group: Group): Observable<Group> {
    console.log('updated');
    return this.http.put<Group>(this.GroupsUrl, group, this.httpOptions);
  }

  deleteGroups(group: Group): Observable<Group> {
    console.log('deleted');
    const url = `api/fullGroups/${group.id}`;
    return this.http.delete<Group>(url, this.httpOptions);
  }

  getUsersAsResource(query: string = ''): Observable<UserAsResource[]> {
    console.log('get');
    return this.http.get<UserAsResource[]>(`api/users${query}`);
  }

  getSpecialisations(): Observable<string[]> {
    console.log('getSpecialisations');
    return this.http.get<string[]>('api/specialisations');
  }
}
