import { Component } from '@angular/core';
import { AuthService } from '../services/auth-service';


@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [],
  templateUrl: './dashboard.html'
})
export class Dashboard {
  constructor(public authService: AuthService) {}
}