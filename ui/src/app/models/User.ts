import * as moment from 'moment';

export interface User {
  userId: number;
  firstName: string;
  lastName: string;
  dateOFBirth: moment.Moment;
  email: string;
  mobilePhone: string;
  createdAt: moment.Moment;
  modifiedAt: moment.Moment;
  userPhoto: any;
}
