import {
  Component,
} from '@angular/core';


import {
  RouterLink,
} from '@angular/router';


import {
  MatListModule,
} from '@angular/material/list';



@Component({

  selector: 'app-sidebar',

  standalone: true,

  imports: [

    RouterLink,

    MatListModule,

  ],

  templateUrl: './sidebar.html',

  styleUrl: './sidebar.scss',

})


export class SidebarComponent {


}