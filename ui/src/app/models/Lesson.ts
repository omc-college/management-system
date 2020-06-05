import * as moment from 'moment';
import {Group} from './Group';
import {User} from './User';
import {Room} from './Room';
import {iSubject} from './Subject';
export interface Lesson {
  readonly id: string;
  subject: iSubject;
<<<<<<< HEAD
  lecturer: User;
  group: Group;
=======
  lecturers: User[];
  groups: Group[];
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
  startAt: moment.Moment;
  endAt: moment.Moment;
  room: Room;
}
