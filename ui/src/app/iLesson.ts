import * as moment from 'moment';
import {Group} from './iGroup';
import {User} from './iUser';
import {Room} from './iRoom';
import {iSubject} from './iSubject';
export interface Lesson {
  readonly id: string;
  subject: iSubject;
  lecturers: User[];
  groups: Group[];
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
