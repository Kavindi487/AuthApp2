import { Component } from '@angular/core';
import { AuthService } from '../services/auth-service';
import {  RouterLink } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './dashboard.html'
})
export class Dashboard {
  constructor(public authService: AuthService) {}
}