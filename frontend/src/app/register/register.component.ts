import {HttpClient} from '@angular/common/http';
import {Component, inject, OnInit} from '@angular/core';
import {FormBuilder, FormControl, FormGroup, ReactiveFormsModule,} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';
import {MatSelectModule} from '@angular/material/select';
import {RegisterService} from '../DAL/register/register.service';
import {CommonModule} from '@angular/common';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSelectModule,
    CommonModule,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
})
export class RegisterComponent implements OnInit {
  industries?: Map<string, number>;
  fb = inject(FormBuilder);
  http = inject(HttpClient);
  register = inject(RegisterService);
  authService = inject(AuthService);
  router = inject(Router);

  form = this.fb.nonNullable.group({
    company: this.fb.nonNullable.group({
      description: this.fb.nonNullable.control(''),
      industry: this.fb.nonNullable.control(''),
      name: this.fb.nonNullable.control(''),
    }),
    owner: this.fb.nonNullable.group({
      department: this.fb.nonNullable.control(''),
      email: this.fb.nonNullable.control(''),
      first_name: this.fb.nonNullable.control(''),
      job_title: this.fb.nonNullable.control(''),
      password: this.fb.nonNullable.control(''),
      second_name: this.fb.nonNullable.control(''),
    }),
  });

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

  get email() {
    return (this.form.get('owner') as FormGroup).get('email') as FormControl;
  }

  get first_name() {
    return (this.form.get('owner') as FormGroup).get(
      'first_name'
    ) as FormControl;
  }

  get second_name() {
    return (this.form.get('owner') as FormGroup).get(
      'second_name'
    ) as FormControl;
  }

  get password() {
    return (this.form.get('owner') as FormGroup).get('password') as FormControl;
  }

  get job_title() {
    return (this.form.get('owner') as FormGroup).get(
      'job_title'
    ) as FormControl;
  }

  get department() {
    return (this.form.get('owner') as FormGroup).get(
      'department'
    ) as FormControl;
  }

  get description() {
    return (this.form.get('company') as FormGroup).get(
      'description'
    ) as FormControl;
  }

  get industry() {
    return (this.form.get('company') as FormGroup).get(
      'industry'
    ) as FormControl;
  }

  get name() {
    return (this.form.get('company') as FormGroup).get('name') as FormControl;
  }

  onSubmit(): void {
    this.register.register(this.form.getRawValue()).subscribe((response) => {
      this.router.navigateByUrl('/');
    });
  }
}
