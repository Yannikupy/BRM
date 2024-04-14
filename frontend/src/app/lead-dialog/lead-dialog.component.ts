import {Component, Inject, inject} from '@angular/core';
import {DragDirective} from "../directives/dragDrop.directive";
import {KeyValuePipe} from "@angular/common";
import {MatButton} from "@angular/material/button";
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogClose,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle
} from "@angular/material/dialog";
import {MatFormField, MatLabel} from "@angular/material/form-field";
import {MatInput} from "@angular/material/input";
import {MatOption} from "@angular/material/autocomplete";
import {MatSelect} from "@angular/material/select";
import {FormBuilder, ReactiveFormsModule} from "@angular/forms";
import {RegisterService} from "../DAL/register/register.service";
import {DalService} from "../DAL/core/dal.service";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-lead-dialog',
  standalone: true,
  imports: [
    DragDirective,
    KeyValuePipe,
    MatButton,
    MatDialogActions,
    MatDialogClose,
    MatDialogContent,
    MatDialogTitle,
    MatFormField,
    MatInput,
    MatLabel,
    MatOption,
    MatSelect,
    ReactiveFormsModule
  ],
  templateUrl: './lead-dialog.component.html',
  styleUrl: './lead-dialog.component.scss'
})
export class LeadDialogComponent {
  fb = inject(FormBuilder)
  register = inject(RegisterService)
  dal = inject(DalService)

  statuses?: Map<string, number>;
  employees?: Map<number, string>;


  adFormGroup = this.fb.group({
    description: this.fb.control(''),
    price: this.fb.control(''),
    responsible: this.fb.control(''),
    status: this.fb.control(''),
    title: this.fb.control('')
  })

  constructor(@Inject(MAT_DIALOG_DATA) public id: number, private _snackBar: MatSnackBar,
              public dialogRef: MatDialogRef<LeadDialogComponent>) {
    this.dal.getLeadById(id).subscribe(value => {
        this.adFormGroup.controls.description.setValue(value.data.description ?? null)
        this.adFormGroup.controls.responsible.setValue(`${value.data.responsible}` ?? null)
        this.adFormGroup.controls.price.setValue(`${value.data.price}` ?? null)
        this.adFormGroup.controls.status.setValue(value.data.status ?? null)
        this.adFormGroup.controls.title.setValue(value.data.title ?? null)
      }
    )

    this.dal.getLeadsStatuses().subscribe(value => {
      this.statuses = new Map();
      for (let key in value.data) {
        this.statuses.set(key, value.data[key]);
      }
    })

    this.dal.getEmployees(100, 0).subscribe(value => {
      this.employees = new Map();
      for (let employee of value.data.employees) {
        this.employees.set(employee.id!, `${employee.second_name} ${employee.first_name}`);
      }
    })
  }

  editLead() {
    this.dal.editLead(this.id!, {
      description: this.adFormGroup.getRawValue().description ?? '',
      price: this.adFormGroup.getRawValue().price ? +this.adFormGroup.getRawValue().price! : 0,
      responsible: +this.adFormGroup.getRawValue().responsible!,
      status: this.adFormGroup.getRawValue().status!,
      title: this.adFormGroup.getRawValue().title ?? ''
    }).subscribe({
      next: (success) => {
        this._snackBar.open('Сделка успешно отредактирована', undefined, {
          duration: 5000,
        });

        this.dialogRef.close()
      },
      error: (error) => {
        this._snackBar.open('Произошла ошибка при редактировании сделки', undefined, {
          duration: 5000,
        });
      }
    })
  }
}
