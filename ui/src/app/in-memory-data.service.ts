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
        startAt: new Date(2020, 4, 17),
        lessonNum: '1',
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
        startAt: new Date(2020, 4, 18),
        lessonNum: '2',
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
        startAt: new Date(2020, 4, 18),
        lessonNum: '4',
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
        startAt: new Date(2020, 4, 19),
        lessonNum: '3',
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
        startAt: new Date(2020, 4, 20),
        lessonNum: '5',
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
        startAt: new Date(2020, 4, 21),
        lessonNum: '6',
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
        startAt: new Date(2020, 4, 22),
        lessonNum: '8',
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
    const timestamp1: string[] = ['1 8.30', '2 9.25', '3 10.30', '4 11.35', '5 12.30', '6 13.25', '7 14.20', '8 15.20'];
    const timestamp2: string[] = [
      '9 14.20',
      '10 15.20',
      '11 16.15',
      '12 17.20',
      '13 18.20',
      '14 19.15',
      '15 20.10',
      '16 21.05',
    ];
    return {lessons, groups, rooms, subjects, users, timestamp1, timestamp2};
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
