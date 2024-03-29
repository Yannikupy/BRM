import {Component, inject} from '@angular/core';
import {MatListModule} from '@angular/material/list';
import {RouterModule} from '@angular/router';
import {DalService} from "../DAL/core/dal.service";
import {AuthService} from "../services/auth.service";

@Component({
  selector: 'app-left-menu',
  standalone: true,
  imports: [MatListModule, RouterModule],
  templateUrl: './left-menu.component.html',
  styleUrl: './left-menu.component.scss',
})
export class LeftMenuComponent {

  dalService = inject(DalService);
  authService = inject(AuthService);
  employeeName: string = ''

  constructor() {
    this.dalService.getEmployeeById(+this.authService.currentUserDataSig()?.
      ["employee-id"]!).subscribe(value => this.employeeName = `${value.data.second_name} ${value.data.first_name}`)
  }
}
