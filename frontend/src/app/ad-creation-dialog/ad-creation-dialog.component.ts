import {Component, inject, OnInit} from '@angular/core';
import {
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogModule,
  MatDialogTitle
} from "@angular/material/dialog";
import {MatFormField, MatFormFieldModule} from "@angular/material/form-field";
import {MatButton} from "@angular/material/button";
import {MatInput} from "@angular/material/input";
import {FormBuilder, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {DragDirective, FileHandle} from "../directives/dragDrop.directive";
import {RegisterService} from "../DAL/register/register.service";
import {KeyValuePipe} from "@angular/common";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";

@Component({
  selector: 'app-ad-creation-dialog',
  standalone: true,
  imports: [
    MatDialogContent,
    MatFormField,
    MatDialogActions,
    MatDialogClose,
    MatButton,
    MatInput,
    MatDialogTitle,
    FormsModule,
    MatDialogModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    DragDirective,
    KeyValuePipe,
    MatOption,
    MatSelect
  ],
  templateUrl: './ad-creation-dialog.component.html',
  styleUrl: './ad-creation-dialog.component.scss'
})
export class AdCreationDialogComponent implements OnInit {
  fb = inject(FormBuilder)
  register = inject(RegisterService)
  industries?: Map<string, number>;

  adFormGroup = this.fb.group({
    image_url: this.fb.control(''),
    industry: this.fb.control(''),
    price: this.fb.control(''),
    text: this.fb.control(''),
    title: this.fb.control('')
  })

  files: FileHandle[] = [];

  filesDropped(files: FileHandle[]): void {
    this.files = files;
  }

  ngOnInit(): void {
    this.register.getIndustries().subscribe({
      next: (success) => {
        this.industries = new Map();
        for (let key in success.data) {
          this.industries.set(key, success.data[key]);
        }
      },
    });
  }
}
