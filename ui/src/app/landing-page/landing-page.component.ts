import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
import {Meta} from '@angular/platform-browser';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.sass'],
})
export class LandingPageComponent implements OnInit {
  constructor(private titleService: Title, private meta: Meta) {}

  ngOnInit(): void {
    this.titleService.setTitle('OMC college');
    this.meta.updateTag({
      name: 'description',
      content:
        'Welcome to Optical and Mechanical College, \
        here you can read information about studying \
        and articles about interesting inventions of our students \
        and other their achievements and watch college`s life.',
    });
    this.meta.updateTag({
      name: 'keywords',
      content:
        'OMC college, student life, student news, education, studying, \
        college, Taras Shevchenko University, programming, economy, physics, journalism',
    });
  }
}
