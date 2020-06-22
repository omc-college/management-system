import {Component, OnInit} from '@angular/core';
import {Title} from '@angular/platform-browser';
import {Meta} from '@angular/platform-browser';

@Component({
  selector: 'app-timetable',
  templateUrl: './timetable.component.html',
  styleUrls: ['./timetable.component.sass'],
})
export class TimetableComponent implements OnInit {
  constructor(private titleService: Title, private meta: Meta) {}

  ngOnInit(): void {
    this.titleService.setTitle('Timetable');
    this.meta.updateTag({
      name: 'description',
      content:
        'Timetabling - one of the part of OMC college`s learning management system. \
        It helps to manage schedule of lessons in college.',
    });
    this.meta.updateTag({
      name: 'keywords',
      content: 'timetable, timetabling, lms, learning management system, OMC college',
    });
  }
}
