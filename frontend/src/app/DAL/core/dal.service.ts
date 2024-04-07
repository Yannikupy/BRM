import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {ContactResponse} from "./model/ContactResponse";
import {environment} from "../../../environments/environment";
import {CompanyResponse} from "./model/CompanyResponse";
import {MainPageResponse} from "./model/MainPageResponse";
import {EmployeeResponse} from "./model/EmployeeResponse";
import {AdListResponse} from "./model/AdListResponse";
import {UpdateContactRequest} from "./model/UpdateContactRequest";
import {AdResponse} from "./model/AdResponse";
import {AddAdRequest} from "./model/AddAdRequest";
import {ResponseResponse} from "./model/ResponseResponse";


@Injectable({
  providedIn: 'root',
})
export class DalService {
  constructor(private _http: HttpClient) {
  }

  getAds(limit: number, offset: number): Observable<AdListResponse> {
    return this._http.get<AdListResponse>(`${environment.coreUrl}/ads?limit=${limit}&offset=${offset}`)
  }

  saveAd(addAdRequest: AddAdRequest) {
    return this._http.post<AdResponse>(`${environment.coreUrl}/ads`, addAdRequest)
  }

  adResponse(id: number): Observable<ResponseResponse> {
    return this._http.post<ResponseResponse>(`${environment.coreUrl}/ads/${id}/response`, null)
  }

  getContacts(limit: number, offset: number): Observable<ContactResponse> {
    return this._http.get<ContactResponse>(`${environment.coreUrl}/contacts?limit=${limit}&offset=${offset}`)
  }

  updateContact(id: number, updateContactRequest: UpdateContactRequest): Observable<ContactResponse> {
    return this._http.put<ContactResponse>(`${environment.coreUrl}/contacts/${id}`, updateContactRequest)
  }

  getCompanyById(id: number): Observable<CompanyResponse> {
    return this._http.get<CompanyResponse>(`${environment.coreUrl}/companies/${id}`)
  }

  getCompanyMainPage(id: number): Observable<MainPageResponse> {
    return this._http.get<MainPageResponse>(`${environment.coreUrl}/companies/${id}/mainpage`)
  }

  getEmployeeById(id: number): Observable<EmployeeResponse> {
    return this._http.get<EmployeeResponse>(`${environment.coreUrl}/employees/${id}`)
  }
}
