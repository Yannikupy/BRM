import {AfterViewInit, Component, inject, ViewChild} from '@angular/core';
import {MatPaginator, MatPaginatorModule} from "@angular/material/paginator";
import {DalService} from "../DAL/core/dal.service";
import {HttpClient} from "@angular/common/http";
import {
  catchError,
  combineLatest,
  concatMap,
  from,
  map,
  Observable,
  of,
  startWith,
  Subscription,
  switchMap,
  takeLast
} from "rxjs";
import {AdListResponse} from "../DAL/core/model/AdListResponse";
import {AdData} from "../DAL/core/model/AdData";
import {MatCardModule} from '@angular/material/card';
import {MatDividerModule} from '@angular/material/divider';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from "@angular/material/icon";
import {NgxSkeletonLoaderModule} from "ngx-skeleton-loader";
import {MatGridListModule} from "@angular/material/grid-list";
import {MatDialog} from "@angular/material/dialog";
import {AdCreationDialogComponent} from "../ad-creation-dialog/ad-creation-dialog.component";
import {FormGroup} from "@angular/forms";
import {AddAdRequest} from "../DAL/core/model/AddAdRequest";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-ads',
  standalone: true,
  imports: [MatPaginatorModule, MatCardModule, MatDividerModule, MatButtonModule, MatIconModule, NgxSkeletonLoaderModule, MatGridListModule],
  templateUrl: './ads.component.html',
  styleUrl: './ads.component.scss'
})
export class AdsComponent implements AfterViewInit {
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  dalService = inject(DalService);
  http = inject(HttpClient);

  ads: AdData[] = []

  subscription = new Subscription()

  imgLoad: boolean = false;

  resultsLength = 0;

  constructor(public dialog: MatDialog, private _snackBar: MatSnackBar) {
    this.loadData(5, 0).subscribe((contacts) => {
      this.ads = contacts.data.ads
      this.resultsLength = contacts.data.amount
    })
  }

  ngAfterViewInit() {
    this.paginator.page
      .pipe(
        startWith({}),
        switchMap(() => {
          return this.loadData(this.paginator.pageSize, this.paginator.pageIndex * this.paginator.pageSize).pipe(catchError(() => of(null)));
        }),
        map(data => {

          if (data === null) {
            return [];
          }

          this.resultsLength = data.data.amount;
          return data.data.ads;
        }),
      )
      .subscribe(data => (this.ads = data));
  }

  loadData(limit: number, offset: number): Observable<AdListResponse> {
    let ads: AdData[] = []

    return this.dalService.getAds(limit, offset).pipe(
      switchMap(value => combineLatest([of(value), from(value.data.ads)])),
      concatMap(([adListResponse, ad]) =>
        combineLatest([of(adListResponse), of(ad), this.dalService.getCompanyById(ad.company_id!)])),
      switchMap(
        ([adListResponse, ad, companyName]) => {
          ad.company_name = companyName.data.name
          ad.imgLoad = false

          ads.push(ad)

          adListResponse.data.ads = ads

          return of(adListResponse)
        }), takeLast(1))


  }

  loadImage(ad: AdData) {
    ad.imgLoad = true
  }

  createAd() {
    this.subscription.add(this.dialog.open(AdCreationDialogComponent, {
      height: '800px',
      width: '800px'
    }).afterClosed().subscribe(value => {
      if (value instanceof FormGroup) {
        this.subscription.add(this.dalService.saveAd(value.getRawValue() as AddAdRequest).subscribe(
          (value) => this.subscription.add(this.loadData(this.paginator.pageSize,
            this.paginator.pageIndex * this.paginator.pageSize)
            .subscribe((contacts) => {
              this.ads = contacts.data.ads
              this.resultsLength = contacts.data.amount
            }))
        ))
      }
    }))
  }

  response(ad: AdData) {
    this.subscription.add(this.dalService.adResponse(ad.id!)
      .subscribe({
        next: (value) =>
          this._snackBar.open('Вы успешно откликнулись на объявление')
      }))
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe()
  }
}
