import {Component, inject, OnDestroy} from '@angular/core';
import {DalService} from "../DAL/core/dal.service";
import {ContactData} from "../DAL/core/model/ContactData";
import {of, Subscription, switchMap} from "rxjs";
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {AsyncPipe} from "@angular/common";
import {MatIconModule} from "@angular/material/icon";
import {MatPaginatorModule} from '@angular/material/paginator';


@Component({
  selector: 'app-contacts',
  standalone: true,
  imports: [MatCardModule, MatDividerModule, MatButtonModule, MatIconModule, MatPaginatorModule, AsyncPipe],
  templateUrl: './contacts.component.html',
  styleUrl: './contacts.component.scss'
})
export class ContactsComponent implements OnDestroy {

  dalService = inject(DalService);

  contacts?: ContactData[]

  subscription = new Subscription()

  constructor() {
    this.subscription.add(this.dalService.getContacts(5, 0).pipe(switchMap(value => {
        for (let contact of value.data) {
          this.subscription.add(this.dalService.getCompanyById(contact.employee?.company_id!).subscribe(
            value => contact.employee!.company_name = value.data.name
          ))
        }
        return of(value)
      })).subscribe(
        value => this.contacts = value.data)
    )
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }

}
