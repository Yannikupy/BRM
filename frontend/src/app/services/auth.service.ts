import { Injectable, signal } from '@angular/core';
import { TokensDataInterface } from '../DAL/login/model/tokens.data.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  currentUserSig = signal<TokensDataInterface | undefined | null>(undefined);
}
