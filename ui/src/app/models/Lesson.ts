import * as moment from 'moment';
import {GroupAsResource} from './GroupAsResource';
import {UserAsResource} from './UserAsResource';
import {Room} from './Room';
import {iSubject} from './Subject';
export interface Lesson {
  readonly id: string;
  subject: iSubject;
  lecturer: UserAsResource;
  group: GroupAsResource;
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
