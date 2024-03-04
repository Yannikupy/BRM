import { ChangeDetectorRef, Component, OnDestroy, inject } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
import { LeftMenuComponent } from '../left-menu/left-menu.component';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MediaMatcher } from '@angular/cdk/layout';
import { RouterOutlet } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { LoginComponent } from '../login/login.component';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-main-page',
  standalone: true,
  imports: [
    CommonModule,
    MatSidenavModule,
    LeftMenuComponent,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    RouterOutlet,
    LoginComponent,
  ],
  templateUrl: './main-page.component.html',
  styleUrl: './main-page.component.scss',
})
export class MainPageComponent {}
