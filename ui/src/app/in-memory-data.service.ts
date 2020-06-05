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
<<<<<<< HEAD
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
=======
        lecturers: [
          {
            id: '11',
            firstName: 'Dmytro',
            lastName: 'Tsapko',
            surname: 'AAAAAAAA',
            role: 'lecturer',
          },
          {
            id: '12',
            firstName: 'Tetyana',
            lastName: 'Lumpova',
            surname: 'Nonamivna',
            role: 'lecturer',
          },
        ],
        groups: [
          {
            id: '131',
            name: 'SD-31',
          },
        ],
        startAt: moment().year(2020).month(5).date(13).hour(8).minute(30),
        endAt: moment().year(2020).month(5).date(13).hour(9).minute(15),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
          {
            id: '12',
            firstName: 'Tetyana',
            lastName: 'Lumpova',
            surname: 'Nonamivna',
            role: 'lecturer',
          },
        ],
        groups: [
          {
            id: '131',
            name: 'SD-31',
          },
          {
            id: '11',
            name: '1A',
          },
        ],
        startAt: moment().year(2020).month(5).date(10).hour(10).minute(30),
        endAt: moment().year(2020).month(5).date(10).hour(11).minute(15),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
          {
            id: '21',
            firstName: 'Borys',
            lastName: 'Gaprindashvili',
            surname: 'Vyachespavovich',
            role: 'lecturer',
          },
        ],
        groups: [
          {
            id: '221',
            name: 'EP-21',
          },
        ],
        startAt: moment().year(2020).month(6).date(11).hour(11).minute(35),
        endAt: moment().year(2020).month(6).date(11).hour(12).minute(20),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
          {
            id: '21',
            firstName: 'Borys',
            lastName: 'Gaprindashvili',
            surname: 'Vyachespavovich',
            role: 'lecturer',
          },
        ],
        groups: [
          {
            id: '222',
            name: 'EP-22',
          },
        ],
        startAt: moment().year(2020).month(5).date(15).hour(14).minute(20),
        endAt: moment().year(2020).month(5).date(15).hour(15).minute(5),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
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
        ],
        groups: [
          {
            id: '221',
            name: 'EP-21',
          },
        ],
        startAt: moment().year(2020).month(4).date(1).hour(13).minute(25),
        endAt: moment().year(2020).month(4).date(1).hour(14).minute(10),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
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
        ],
        groups: [
          {
            id: '311',
            name: 'OO-11',
          },
        ],
        startAt: moment().year(2021).month(2).date(9).hour(10).minute(30),
        endAt: moment().year(2021).month(2).date(9).hour(11).minute(25),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
        lecturers: [
          {
            id: '32',
            firstName: 'Tetyana',
            lastName: 'Garivna',
            surname: 'Noname',
            role: 'director',
          },
          {
            id: '11',
            firstName: 'Dmytro',
            lastName: 'Tsapko',
            surname: 'AAAAAAAA',
            role: 'lecturer',
          },
        ],
        groups: [
          {
            id: '311',
            name: 'OO-11',
          },
        ],
        startAt: moment().year(2020).month(12).date(12).hour(8).minute(30),
        endAt: moment().year(2020).month(12).date(12).hour(9).minute(15),
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
=======
        id: '11',
        firstName: 'Dmytro',
        lastName: 'Tsapko',
        surname: 'AAAAAAAA',
        role: 'lecturer',
      },
      {
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
<<<<<<< HEAD
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
=======
      '8.30-9.15',
      '9.25-10.10',
      '10.30-11.15',
      '11.35-12.20',
      '12.30-13.15',
      '13.25-14.10',
      '14.20-15.05',
      '15.20-16.05',
    ];
    const timestamp2: string[] = [
      '14.20-15.05',
      '15.20-16.05',
      '16.15-17.00',
      '17.20-18.05',
      '18.20-19.05',
      '19.15-20.00',
      '20.10-20.55',
      '21.05-21.50',
    ];
    return {lessons, groups, rooms, subjects, users, timestamp1, timestamp2};
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
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
