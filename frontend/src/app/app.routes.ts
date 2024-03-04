import { Routes } from '@angular/router';
import { MainPageComponent } from './main-page/main-page.component';
import { ContactsComponent } from './contacts/contacts.component';
import { LoginComponent } from './login/login.component';
import { authGuard } from './guards/auth.guard';
import { AppComponent } from './app.component';

export const routes: Routes = [
  {
    path: 'contacts',
    component: ContactsComponent,
    canActivate: [authGuard],
  },
  { path: 'login', component: LoginComponent },
  { path: '', component: AppComponent, canActivate: [authGuard] },
];
