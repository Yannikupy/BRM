import {Component} from '@angular/core';
import {MatListModule} from '@angular/material/list';
import {RouterModule} from '@angular/router';

@Component({
  selector: 'app-left-menu',
  standalone: true,
  imports: [MatListModule, RouterModule],
  templateUrl: './left-menu.component.html',
  styleUrl: './left-menu.component.scss',
})
export class LeftMenuComponent {
}
