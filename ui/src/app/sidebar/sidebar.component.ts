import { Component, OnInit, Injectable } from '@angular/core';
import { Error } from '../error';

@Component({
   selector: 'app-sidebar',
   templateUrl: './sidebar.component.html',
   styleUrls: ['./sidebar.component.sass'],
})
@Injectable({
   providedIn: 'root',
})
export class SidebarComponent implements OnInit {
   error: Error = {
      id: '55',
      code: 505,
      message: 'Here will be errors or other messages',
   };
   constructor() {}

   ngOnInit(): void {}
}
