import {Injectable} from '@angular/core';
import {InMemoryDbService} from 'angular-in-memory-web-api';
import * as moment from 'moment';

import {Lesson} from './models/Lesson';
import {Group} from './models/Group';
import {Room} from './models/Room';
import {iSubject} from './models/Subject';
import {User} from './models/User';

@Injectable({
  providedIn: 'root',
})
export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    const lessons: Lesson[] = [
      {
        id: '1',
        subject: {
          id: '11',
          nameOfSubject: 'Programming',
        },
        lecturer: {
          id: '21',
          firstName: 'Borys',
          lastName: 'Gaprindashvili',
          surname: 'Vyachespavovich',
          role: 'lecturer',
        },
        group: {
          id: '131',
          name: 'SD-31',
        },
        startAt: moment([2020, 4, 18, 8, 30]),
        endAt: moment([2020, 4, 18, 9, 15]),
        room: {
          id: '2011',
          room: '201A',
        },
      },
      {
        id: '2',
        subject: {
          id: '12',
          nameOfSubject: 'Math',
        },
        lecturer: {
          id: '12',
          firstName: 'Tetyana',
          lastName: 'Lumpova',
          surname: 'Nonamivna',
          role: 'lecturer',
        },
        group: {
          id: '131',
          name: 'SD-31',
        },
        startAt: moment([2020, 4, 18, 10, 30]),
        endAt: moment([2020, 4, 18, 12, 15]),
        room: {
          id: '2011',
          room: '201A',
        },
      },
      {
        id: '3',
        subject: {
          id: '13',
          nameOfSubject: 'Physics',
        },
        lecturer: {
          id: '21',
          firstName: 'Borys',
          lastName: 'Gaprindashvili',
          surname: 'Vyachespavovich',
          role: 'lecturer',
        },
        group: {
          id: '221',
          name: 'EP-21',
        },
        startAt: moment([2020, 4, 19, 11, 35]),
        endAt: moment([2020, 4, 19, 13, 10]),
        room: {
          id: '2012',
          room: '201B',
        },
      },
      {
        id: '4',
        subject: {
          id: '14',
          nameOfSubject: 'English',
        },
        lecturer: {
          id: '21',
          firstName: 'Borys',
          lastName: 'Gaprindashvili',
          surname: 'Vyachespavovich',
          role: 'lecturer',
        },
        group: {
          id: '222',
          name: 'EP-22',
        },
        startAt: moment([2020, 4, 19, 13, 25]),
        endAt: moment([2020, 4, 19, 14, 20]),
        room: {
          id: '45',
          room: '45',
        },
      },
      {
        id: '5',
        subject: {
          id: '14',
          nameOfSubject: 'English',
        },
        lecturer: {
          id: '21',
          firstName: 'Borys',
          lastName: 'Gaprindashvili',
          surname: 'Vyachespavovich',
          role: 'lecturer',
        },
        group: {
          id: '221',
          name: 'EP-21',
        },
        startAt: moment([2020, 4, 20, 14]),
        endAt: moment([2020, 4, 20, 15, 10]),
        room: {
          id: '31',
          room: '3A',
        },
      },
      {
        id: '6',
        subject: {
          id: '15',
          nameOfSubject: 'Ukrainian language',
        },
        lecturer: {
          id: '21',
          firstName: 'Borys',
          lastName: 'Gaprindashvili',
          surname: 'Vyachespavovich',
          role: 'lecturer',
        },
        group: {
          id: '311',
          name: 'OO-11',
        },
        startAt: moment([2020, 4, 21, 8, 30]),
        endAt: moment([2020, 4, 21, 9, 15]),
        room: {
          id: '202',
          room: '20B',
        },
      },
      {
        id: '7',
        subject: {
          id: '11',
          nameOfSubject: 'History',
        },
        lecturer: {
          id: '32',
          firstName: 'Tetyana',
          lastName: 'Garivna',
          surname: 'Noname',
          role: 'director',
        },
        group: {
          id: '311',
          name: 'OO-11',
        },
        startAt: moment([2020, 4, 22, 9, 30]),
        endAt: moment([2020, 4, 22, 10, 15]),
        room: {
          id: '12',
          room: '12',
        },
      },
    ];
    const users: User[] = [
      {
        id: '1',
        firstName: 'Andriy',
        lastName: 'Vashchuk',
        surname: 'hzhzhzhz',
        role: 'student',
      },
      {
        id: '12',
        firstName: 'Tetyana',
        lastName: 'Lumpova',
        surname: 'Nonamivna',
        role: 'lecturer',
      },
      {
        id: '21',
        firstName: 'Borys',
        lastName: 'Gaprindashvili',
        surname: 'Vyachespavovich',
        role: 'lecturer',
      },
      {
        id: '32',
        firstName: 'Tetyana',
        lastName: 'Garivna',
        surname: 'Noname',
        role: 'director',
      },
      {
        id: '332',
        firstName: 'Vasya',
        lastName: 'Parent',
        surname: 'Pupkin',
        role: 'parent',
      },
    ];
    const groups: Group[] = [
      {
        id: '131',
        name: 'SD-31',
      },
      {
        id: '132',
        name: 'SD-32',
      },
      {
        id: '221',
        name: 'EP-21',
      },
      {
        id: '222',
        name: 'EP-22',
      },
      {
        id: '311',
        name: 'OO-11',
      },
      {
        id: '312',
        name: 'OO-12',
      },
    ];
    const rooms: Room[] = [
      {
        id: '2011',
        room: '201A',
      },
      {
        id: '2012',
        room: '201B',
      },
      {
        id: '45',
        room: '45',
      },
      {
        id: '31',
        room: '3A',
      },
      {
        id: '202',
        room: '20B',
      },
      {
        id: '12',
        room: '12',
      },
      {
        id: '711',
        room: '71A',
      },
      {
        id: '88',
        room: '88',
      },
      {
        id: '91',
        room: '91',
      },
      {
        id: '4',
        room: '4',
      },
    ];
    const subjects: iSubject[] = [
      {
        id: '11',
        nameOfSubject: 'Programming',
      },
      {
        id: '12',
        nameOfSubject: 'Math',
      },
      {
        id: '13',
        nameOfSubject: 'Physics',
      },
      {
        id: '14',
        nameOfSubject: 'English',
      },
      {
        id: '15',
        nameOfSubject: 'Ukrainian language',
      },
      {
        id: '16',
        nameOfSubject: 'History',
      },
      {
        id: '17',
        nameOfSubject: 'Literature',
      },
    ];
    const timestamp1: string[] = [
      '8',
      '9',
      '10',
      '11',
      '12',
      '13',
      '14',
      '15',
      '16',
      '17',
      '18',
      '19',
      '20',
      '21',
      '22',
      '23',
    ];

    return {lessons, groups, rooms, subjects, users, timestamp1};
  }

  // Overrides the genId method to ensure that a hero always has an id.
  // If the heroes array is empty,
  // the method below returns the initial number (11).
  // if the heroes array is not empty, the method below returns the highest
  // hero id + 1.
  genId(items: Lesson[]): string {
    return items.length > 0 ? String(Math.max(...items.map(lesson => +lesson.id)) + 1) : '0';
  }
}
