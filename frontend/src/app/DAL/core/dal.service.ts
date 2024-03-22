import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {ContactResponse} from "./model/ContactResponse";
import {environment} from "../../../environments/environment";


@Injectable({
  providedIn: 'root',
})
export class DalService {
  constructor(private _http: HttpClient) {
  }

  getContacts(limit: number, offset: number): Observable<ContactResponse> {
    return this._http.get<ContactResponse>(`${environment.coreUrl}/contacts?limit=${limit}&offset=${offset}`)
  }
}
