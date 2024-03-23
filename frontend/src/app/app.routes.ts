import {Routes} from '@angular/router';
import {ContactsComponent} from './contacts/contacts.component';
import {LoginComponent} from './login/login.component';
import {authGuard} from './guards/auth.guard';
import {RegisterComponent} from './register/register.component';
import {CompanyComponent} from "./company/company.component";

export const routes: Routes = [
  {
    path: 'contacts',
    component: ContactsComponent,
    canActivate: [authGuard],
  },
  {path: 'company', component: CompanyComponent, canActivate: [authGuard]},
  {path: 'login', component: LoginComponent},
  {path: 'sign-up', component: RegisterComponent},
];
