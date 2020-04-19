import {Lesson} from './iLesson';
import * as moment from 'moment';
export const LESSONS: Lesson[] = [
  {
    id: '1',
    subject: {
      id: '11',
      nameOfSubject: 'Programming',
    },
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
    room: {
      id: '2011',
      room: '201A',
    },
  },
  {
    id: '5',
    subject: {
      id: '12',
      nameOfSubject: 'Math',
    },
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
    room: {
      id: '2011',
      room: '201A',
    },
  },
  {
    id: '7',
    subject: {
      id: '13',
      nameOfSubject: 'Physics',
    },
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
    room: {
      id: '2012',
      room: '201B',
    },
  },
  {
    id: '225',
    subject: {
      id: '14',
      nameOfSubject: 'English',
    },
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
    room: {
      id: '45',
      room: '45',
    },
  },
  {
    id: '11',
    subject: {
      id: '14',
      nameOfSubject: 'English',
    },
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
    room: {
      id: '31',
      room: '3A',
    },
  },
  {
    id: '45',
    subject: {
      id: '15',
      nameOfSubject: 'Ukrainian language',
    },
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
    room: {
      id: '202',
      room: '20B',
    },
  },
  {
    id: '1',
    subject: {
      id: '11',
      nameOfSubject: 'History',
    },
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
    room: {
      id: '12',
      room: '12',
    },
  },
];
