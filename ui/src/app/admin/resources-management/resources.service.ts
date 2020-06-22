import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

import {Resource} from '../../models/Resource';

@Injectable({
  providedIn: 'root',
})
export class ResourcesService {
  private resourcesUrl = 'api/resources';
  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'}),
  };

  constructor(private http: HttpClient) {}

  getResources(): Observable<Resource[]> {
    console.log('get');
    return this.http.get<Resource[]>(this.resourcesUrl);
  }

  addResources(resource: Resource): Observable<Resource> {
    console.log('added');
    return this.http.post<Resource>(this.resourcesUrl, resource, this.httpOptions);
  }

  deleteResources(resource: Resource): Observable<Resource> {
    console.log('deleted');
    const url = `api/Resources/${resource.resourceId}`;
    return this.http.delete<Resource>(url, this.httpOptions);
  }

  updateResources(resource): Observable<Resource> {
    console.log('updated');
    return this.http.put<Resource>(this.resourcesUrl, resource, this.httpOptions);
  }
}
