import {Component, inject} from '@angular/core';
import {DalService} from "../DAL/core/dal.service";
import {ContactData} from "../DAL/core/model/ContactData";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-contacts',
  standalone: true,
  imports: [],
  templateUrl: './contacts.component.html',
  styleUrl: './contacts.component.scss'
})
export class ContactsComponent {

  dalService = inject(DalService);

  contacts?: ContactData[]

  subscription = new Subscription()

  constructor() {
    this.subscription.add(this.dalService.getContacts(0, 1).subscribe(value => this.contacts = value.data))
  }

}
