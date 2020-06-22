import {Injectable} from '@angular/core';
import {InMemoryDbService} from 'angular-in-memory-web-api';
import * as moment from 'moment';

import {Lesson} from '../models/Lesson';
import {GroupAsResource} from '../models/GroupAsResource';
import {Room} from '../models/Room';
import {iSubject} from '../models/Subject';
import {UserAsResource} from '../models/UserAsResource';
import {Role, FeatureEntry, Endpoint} from '../models/role';
import {Resource} from '../models/Resource';
import {User} from '../models/User';

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
    const fullUsers: User[] = [
      {
        userId: 9,
        firstName: 'Name',
        lastName: 'Lastname',
        surname: 'Surname',
        dateOFBirth: moment('08-06-2020', 'DD-MM-YYYY'),
        email: 'www.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['admin', 'director', 'teacher'],
        verified: false,
        userPhoto: 'photo',
      },
      {
        userId: 1,
        firstName: 'Name2',
        lastName: 'Lastnam2',
        surname: 'Surnam2',
        dateOFBirth: moment('07-06-2020', 'DD-MM-YYYY'),
        email: 'sw.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['student'],
        verified: true,
        userPhoto: 'photo2',
      },
      {
        userId: 2,
        firstName: 'Name3',
        lastName: 'Lastnam3',
        surname: 'Surnam3',
        dateOFBirth: moment('09-05-2020', 'DD-MM-YYYY'),
        email: 'awww.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('07-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['teacher', 'parent'],
        verified: true,
        userPhoto: 'photo3',
      },
      {
        userId: 0,
        firstName: 'Name',
        lastName: 'Lastname',
        surname: 'Surname',
        dateOFBirth: moment('08-06-2020', 'DD-MM-YYYY'),
        email: 'www.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['admin', 'director', 'teacher'],
        verified: false,
        userPhoto: 'photo',
      },
      {
        userId: 1,
        firstName: 'Name2',
        lastName: 'Lastnam2',
        surname: 'Surnam2',
        dateOFBirth: moment('07-06-2020', 'DD-MM-YYYY'),
        email: 'sw.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['student'],
        verified: true,
        userPhoto: 'photo2',
      },
      {
        userId: 2,
        firstName: 'Name3',
        lastName: 'Lastnam3',
        surname: 'Surnam3',
        dateOFBirth: moment('09-05-2020', 'DD-MM-YYYY'),
        email: 'awww.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('07-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['teacher', 'parent'],
        verified: true,
        userPhoto: 'photo3',
      },
      {
        userId: 0,
        firstName: 'Name',
        lastName: 'Lastname',
        surname: 'Surname',
        dateOFBirth: moment('08-06-2020', 'DD-MM-YYYY'),
        email: 'www.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['admin', 'director', 'teacher'],
        verified: false,
        userPhoto: 'photo',
      },
      {
        userId: 1,
        firstName: 'Name2',
        lastName: 'Lastnam2',
        surname: 'Surnam2',
        dateOFBirth: moment('07-06-2020', 'DD-MM-YYYY'),
        email: 'sw.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['student'],
        verified: true,
        userPhoto: 'photo2',
      },
      {
        userId: 2,
        firstName: 'Name3',
        lastName: 'Lastnam3',
        surname: 'Surnam3',
        dateOFBirth: moment('09-05-2020', 'DD-MM-YYYY'),
        email: 'awww.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('07-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['teacher', 'parent'],
        verified: true,
        userPhoto: 'photo3',
      },
      {
        userId: 0,
        firstName: 'Name',
        lastName: 'Lastname',
        surname: 'Surname',
        dateOFBirth: moment('08-06-2020', 'DD-MM-YYYY'),
        email: 'www.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['admin', 'director', 'teacher'],
        verified: false,
        userPhoto: 'photo',
      },
      {
        userId: 1,
        firstName: 'Name2',
        lastName: 'Lastnam2',
        surname: 'Surnam2',
        dateOFBirth: moment('07-06-2020', 'DD-MM-YYYY'),
        email: 'sw.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('08-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['student'],
        verified: true,
        userPhoto: 'photo2',
      },
      {
        userId: 2,
        firstName: 'Name3',
        lastName: 'Lastnam3',
        surname: 'Surnam3',
        dateOFBirth: moment('09-05-2020', 'DD-MM-YYYY'),
        email: 'awww.ex.dmail.com',
        mobilePhone: '801475918',
        createdAt: moment('07-06-2020', 'DD-MM-YYYY'),
        modifiedAt: moment(),
        roles: ['teacher', 'parent'],
        verified: true,
        userPhoto: 'photo3',
      },
    ];
    const users: UserAsResource[] = [
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
    const groups: GroupAsResource[] = [
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
    // ------------------------------------------Roles--------------------------------------------------
    const roles: Role[] = [
      {
        id: 0,
        name: 'Administrator',
        entries: [
          {
            id: 1,
            name: 'roles managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 2,
        name: 'Manager',
        entries: [
          {
            id: 1,
            name: 'smth managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
              {
                id: 3,
                name: 'post role',
                path: '/www/nnn',
                method: 'delete',
              },
            ],
          },
          {
            id: 2,
            name: 'smth222 managenent',
            description:
              'ffffffffffffff vvdfs ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 3,
        name: 'Lecturer',
        entries: [
          {
            id: 7,
            name: 'roles managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 4,
        name: 'Student',
        entries: [
          {
            id: 1,
            name: 'smth managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
              {
                id: 3,
                name: 'post role',
                path: '/www/nnn',
                method: 'delete',
              },
            ],
          },
          {
            id: 2,
            name: 'smth222 managenent',
            description: 'ffffffffffffff df dfffdfsd sfgsdfsdf f sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 0,
        name: 'Administrator',
        entries: [
          {
            id: 1,
            name: 'roles managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 2,
        name: 'Manager',
        entries: [
          {
            id: 1,
            name: 'smth managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
              {
                id: 3,
                name: 'post role',
                path: '/www/nnn',
                method: 'delete',
              },
            ],
          },
          {
            id: 2,
            name: 'smth222 managenent',
            description: 'fffffffffdddd f dffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 3,
        name: 'Lecturer',
        entries: [
          {
            id: 1,
            name: 'roles managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
      {
        id: 4,
        name: 'Student',
        entries: [
          {
            id: 1,
            name: 'smth managenent',
            description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
              {
                id: 3,
                name: 'post role',
                path: '/www/nnn',
                method: 'delete',
              },
            ],
          },
          {
            id: 2,
            name: 'smth222 managenent',
            description: 'fffffffffffffdd f df dfffdfsd sfgsdfsd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
            endpoints: [
              {
                id: 2,
                name: 'Del role',
                path: '/www/exe',
                method: 'post',
              },
            ],
          },
        ],
      },
    ];
    const endpoints: Endpoint[] = [
      {
        id: 1,
        name: 'Delete Lesson',
        path: '/timetable',
        method: 'delete',
      },
      {
        id: 2,
        name: 'Add Lesson',
        path: '/timetable',
        method: 'post',
      },
      {
        id: 3,
        name: 'Update Lesson',
        path: '/timetable',
        method: 'put',
      },
      {
        id: 4,
        name: 'Get Lesson',
        path: '/timetable',
        method: 'get',
      },
      {
        id: 5,
        name: 'Delete lecturer',
        path: '/admin',
        method: 'delete',
      },
      {
        id: 6,
        name: 'Add lecturer',
        path: '/admin',
        method: 'delete',
      },
      {
        id: 7,
        name: 'Update lecturer',
        path: '/admin',
        method: 'put',
      },
      {
        id: 8,
        name: 'Get lecturer',
        path: '/admin',
        method: 'get',
      },
    ];
    const featureEntry: FeatureEntry[] = [
      {
        id: 7,
        name: 'roles managenent',
        description: 'ffffffffffffffff dddddddd f df dfffdfsd sfgsdfsdf sdfsdfsd vvvsvsd vvdfs',
        endpoints: [
          {
            id: 2,
            name: 'Del role',
            path: '/www/exe',
            method: 'post',
          },
        ],
      },
      {
        id: 8,
        name: 'Full access to lessons',
        description: 'c r u d',
        endpoints: [
          {
            id: 1,
            name: 'Delete Lesson',
            path: '/timetable',
            method: 'delete',
          },
          {
            id: 2,
            name: 'Add Lesson',
            path: '/timetable',
            method: 'post',
          },
          {
            id: 3,
            name: 'Update Lesson',
            path: '/timetable',
            method: 'put',
          },
          {
            id: 4,
            name: 'Get Lesson',
            path: '/timetable',
            method: 'get',
          },
        ],
      },
      {
        id: 9,
        name: 'Full acess to lecturers',
        description: 'c r u d',
        endpoints: [
          {
            id: 5,
            name: 'Delete lecturer',
            path: '/admin',
            method: 'delete',
          },
          {
            id: 6,
            name: 'Add lecturer',
            path: '/admin',
            method: 'delete',
          },
          {
            id: 7,
            name: 'Update lecturer',
            path: '/admin',
            method: 'put',
          },
          {
            id: 8,
            name: 'Get lecturer',
            path: '/admin',
            method: 'get',
          },
        ],
      },
    ];
    const resources: Resource[] = [
      {
        resourceId: 0,
        resourceName: 'Projector',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 1,
        resourceName: 'Projector 2',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 2,
        resourceName: 'Projector 3',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 3,
        resourceName: 'Projector 4',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 4,
        resourceName: 'Projector 5',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 5,
        resourceName: 'Projector 6',
        resourceDescription: 'Required when you are needed to show presentation',
      },
      {
        resourceId: 6,
        resourceName: 'Projector 7',
        resourceDescription: 'Required when you are needed to show presentation',
      },
    ];
    return {lessons, groups, rooms, subjects, users, timestamp1, roles, endpoints, featureEntry, resources, fullUsers};
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
