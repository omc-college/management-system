import * as moment from 'moment';
import {Group} from './Group';
import {User} from './User';
import {Room} from './Room';
import {iSubject} from './Subject';
export interface Lesson {
  readonly id: string;
  subject: iSubject;
  lecturers: User[];
  groups: Group[];
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
