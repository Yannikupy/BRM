import {Component, inject} from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {DalService} from "../DAL/core/dal.service";
import {MainPageResponse} from "../DAL/core/model/MainPageResponse";

@Component({
  selector: 'app-company',
  standalone: true,
  imports: [MatCardModule],
  templateUrl: './company.component.html',
  styleUrl: './company.component.scss'
})
export class CompanyComponent {

  dalService = inject(DalService);
  mainPage?: MainPageResponse

  constructor() {
    this.dalService.getCompanyMainPage(1).subscribe(
      value => this.mainPage = value
    )
  }
}
