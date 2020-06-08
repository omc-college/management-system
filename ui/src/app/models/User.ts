import * as moment from 'moment';

export interface User {
  userId: number;
  firstName: string;
  lastName: string;
  surname: string;
  dateOFBirth: moment.Moment;
  email: string;
  mobilePhone: string;
  createdAt: moment.Moment;
  modifiedAt: moment.Moment;
  roles: string[];
  verified: boolean;
  userPhoto: any;
}
