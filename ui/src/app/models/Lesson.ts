import * as moment from 'moment';
import {GroupAsResource} from './GroupAsResource';
import {UserAsResource} from './UserAsResource';
import {Room} from './Room';
import {SubjectInterface} from './Subject';
export interface Lesson {
  readonly id: string;
  subject: SubjectInterface;
  lecturer: UserAsResource;
  group: GroupAsResource;
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
