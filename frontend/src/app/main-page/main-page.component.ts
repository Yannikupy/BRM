import { Component } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
import { LeftMenuComponent } from '../left-menu/left-menu.component';
import { RightMenuComponent } from '../right-menu/right-menu.component';

@Component({
  selector: 'app-main-page',
  standalone: true,
  imports: [MatSidenavModule, LeftMenuComponent, RightMenuComponent],
  templateUrl: './main-page.component.html',
  styleUrl: './main-page.component.scss',
})
export class MainPageComponent {}
