import {Routes} from '@angular/router';
import {ContactsComponent} from './contacts/contacts.component';
import {LoginComponent} from './login/login.component';
import {authGuard} from './guards/auth.guard';
import {RegisterComponent} from './register/register.component';
import {CompanyComponent} from "./company/company.component";
import {AdsComponent} from "./ads/ads.component";
import {SettingsComponent} from "./settings/settings.component";
import {LeadsComponent} from "./leads/leads.component";

export const routes: Routes = [
  {path: '', redirectTo: 'company', pathMatch: "full"},
  {
    path: 'contacts',
    component: ContactsComponent,
    canActivate: [authGuard],
  },
  {path: 'leads', component: LeadsComponent, canActivate: [authGuard]},
  {path: 'settings', component: SettingsComponent, canActivate: [authGuard]},
  {path: 'company', component: CompanyComponent, canActivate: [authGuard]},
  {path: 'ads', component: AdsComponent, canActivate: [authGuard]},
  {path: 'login', component: LoginComponent},
  {path: 'sign-up', component: RegisterComponent},
];
