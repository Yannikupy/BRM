import {
  ChangeDetectorRef,
  Component,
  OnDestroy,
  OnInit,
  inject,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterOutlet } from '@angular/router';
import { MediaMatcher } from '@angular/cdk/layout';
import { AuthService } from './services/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { LeftMenuComponent } from './left-menu/left-menu.component';
import { LoginComponent } from './login/login.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    MatSidenavModule,
    LeftMenuComponent,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    LoginComponent,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit, OnDestroy {
  title = 'BRM';
  mobileQuery: MediaQueryList;
  router = inject(Router);

  authService = inject(AuthService);

  private _mobileQueryListener: () => void;

  constructor(changeDetectorRef: ChangeDetectorRef, media: MediaMatcher) {
    this.mobileQuery = media.matchMedia('(max-width: 600px)');
    this._mobileQueryListener = () => changeDetectorRef.detectChanges();
    this.mobileQuery.addListener(this._mobileQueryListener);
  }

  ngOnInit(): void {
    const token = localStorage.getItem('token');
    if (token && token != '') {
      this.authService.currentUserSig.set({ access: token, refresh: '' });
      this.router.navigateByUrl('/');
    }
  }

  ngOnDestroy(): void {
    this.mobileQuery.removeListener(this._mobileQueryListener);
  }

  logout(): void {
    localStorage.setItem('token', '');
    this.authService.currentUserSig.set(null);
    this.router.navigateByUrl('/login');
  }
}
