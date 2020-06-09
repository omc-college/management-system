import {UserAsResource} from './UserAsResource';
export interface Group {
  id: number;
  specialisation: string;
  yearOfEducation: number;
  groupNumber: number;
  curator: UserAsResource;
  students: UserAsResource[];
}
