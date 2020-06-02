import {User} from './User';
export interface Group {
  id: number;
  specialisation: string;
  yearOfEducation: number;
  groupNumber: number;
  curator: User;
  students: User[];
}
