import { Component, OnInit, Injectable } from '@angular/core';

@Component({
   selector: 'app-header',
   templateUrl: './header.component.html',
   styleUrls: ['./header.component.sass'],
})
@Injectable({
   providedIn: 'root',
})
export class HeaderComponent implements OnInit {
   constructor() {}

   ngOnInit(): void {}
}
