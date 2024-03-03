import { Routes } from '@angular/router';
import { MainPageComponent } from './main-page/main-page.component';
import { ContactsComponent } from './contacts/contacts.component';
import { LoginComponent } from './login/login.component';

export const routes: Routes = [
  {
    path: 'contacts',
    component: ContactsComponent,
  },
  { path: 'login', component: LoginComponent },
];
