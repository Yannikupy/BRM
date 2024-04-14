import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {MatIconModule} from "@angular/material/icon";
import {AuthService} from "../services/auth.service";
import {Router} from "@angular/router";
import {MatButtonModule} from "@angular/material/button";
import {Subscription} from "rxjs";
import {DalService} from "../DAL/core/dal.service";

@Component({
  selector: 'app-toolbar',
  standalone: true,
  imports: [MatIconModule, MatButtonModule],
  templateUrl: './toolbar.component.html',
  styleUrl: './toolbar.component.scss'
})
export class ToolbarComponent implements OnInit, OnDestroy {
  @Input() snav: any

  authService = inject(AuthService);
  dalService = inject(DalService);

  subscription: Subscription = new Subscription()
  router = inject(Router);
  companyName: string = ''

  ngOnInit(): void {
    this.subscription.add(this.dalService.getCompanyById(+this.authService.currentUserDataSig()?.
      ["company-id"]!).subscribe((
      value => this.companyName = value.data.name!
    )))
  }

  ngOnDestroy() {
    this.subscription.unsubscribe()
  }

  logout(): void {
    localStorage.setItem('token', '');
    this.authService.currentUserSig.set(null);
    this.authService.currentUserDataSig.set(null);
    this.snav.close();
    this.router.navigateByUrl('/login');
  }
}
