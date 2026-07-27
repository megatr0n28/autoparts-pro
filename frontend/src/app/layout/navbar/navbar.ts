import {
  Component,
  OnInit,
} from '@angular/core';


import {
  Router,
} from '@angular/router';


import {
  MatToolbarModule,
} from '@angular/material/toolbar';


import {
  MatButtonModule,
} from '@angular/material/button';


import {
  AuthService,
} from '../../core/services/auth.service';



@Component({

  selector: 'app-navbar',

  standalone: true,

  imports: [

    MatToolbarModule,

    MatButtonModule,

  ],

  templateUrl: './navbar.html',

  styleUrl: './navbar.scss',

})


export class NavbarComponent
implements OnInit {


  userId = '';

  role = '';



  constructor(

    private auth: AuthService,

    private router: Router,

  ) {}



  ngOnInit(): void {


    this.auth
      .getCurrentUser()
      .subscribe({

        next: user => {

          this.userId =
            user.user_id;

          this.role =
            user.role;

        },


        error: err => {

          console.error(
            "Failed loading user",
            err
          );

        }

      });


  }



  logout(): void {


    this.auth.logout();


    this.router.navigate([
      '/login'
    ]);


  }


}