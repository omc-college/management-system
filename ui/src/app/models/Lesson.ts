import * as moment from 'moment';
import {Group} from './Group';
import {User} from './User';
import {Room} from './Room';
import {iSubject} from './Subject';
export interface Lesson {
  readonly id: string;
  subject: iSubject;
  lecturer: User;
  group: Group;
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
