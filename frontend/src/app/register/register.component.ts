import { HttpClient } from '@angular/common/http';
import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
})
export class RegisterComponent {
  fb = inject(FormBuilder);
  http = inject(HttpClient);
  authService = inject(AuthService);
  router = inject(Router);

  form = this.fb.nonNullable.group({
    company: this.fb.nonNullable.group({
      description: this.fb.control(''),
      industry: this.fb.control(''),
      name: this.fb.control(''),
    }),
    owner: this.fb.nonNullable.group({
      department: this.fb.control(''),
      email: this.fb.control(''),
      first_name: this.fb.control(''),
      job_title: this.fb.control(''),
      password: this.fb.control(''),
      second_name: this.fb.control(''),
    }),
  });

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

  onSubmit(): void {
    /*     this.loginService.login(this.form.getRawValue()).subscribe((response) => {
      console.log('response', response);
      localStorage.setItem('token', response.data.access);
      this.authService.currentUserSig.set(response.data);
      this.router.navigateByUrl('/');
    }); */
  }
}
