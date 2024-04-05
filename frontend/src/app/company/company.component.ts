import {Component, inject} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {DalService} from "../DAL/core/dal.service";
import {MainPageResponse} from "../DAL/core/model/MainPageResponse";
import {AuthService} from "../services/auth.service";

@Component({
  selector: 'app-company',
  standalone: true,
  imports: [MatCardModule],
  templateUrl: './company.component.html',
  styleUrl: './company.component.scss'
})
export class CompanyComponent {

  dalService = inject(DalService);
  authService = inject(AuthService);
  mainPage?: MainPageResponse

  constructor() {
    this.dalService.getCompanyMainPage(+this.authService.currentUserDataSig()?.
      ["company-id"]!).subscribe(
      value => this.mainPage = value
    )
  }
}
