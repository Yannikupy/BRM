import {Component, Inject, OnInit, Renderer2} from '@angular/core';
import {MatSlideToggle} from "@angular/material/slide-toggle";
import {FormControl, ReactiveFormsModule} from "@angular/forms";
import {DOCUMENT} from "@angular/common";
import {MatCard, MatCardContent, MatCardHeader, MatCardModule} from "@angular/material/card";

@Component({
  selector: 'app-settings',
  standalone: true,
  imports: [
    MatSlideToggle,
    ReactiveFormsModule,
    MatCard,
    MatCardHeader,
    MatCardContent,
    MatCardModule
  ],
  templateUrl: './settings.component.html',
  styleUrl: './settings.component.scss'
})
export class SettingsComponent implements OnInit {

  switchTheme = new FormControl(false)

  darkClass = 'theme-dark'
  lightClass = 'theme-light'

  constructor(@Inject(DOCUMENT) private document: Document, private renderer: Renderer2) {
  }

  ngOnInit() {
    this.switchTheme.setValue(this.document.body.classList.value !== this.lightClass)
    this.switchTheme.valueChanges.subscribe((currentMode) => {
      const className = currentMode ? this.darkClass : this.lightClass
      this.renderer.setAttribute(this.document.body, 'class', className)
    })
  }

}
