interface Group {
   readonly id: string;
   name: string;
}
interface Lecturer {
   readonly id: string;
   firstName: string;
   lastName: string;
   surname: string;
}
export interface Lesson {
   readonly id: string;
   subject: {
      readonly id: string;
      nameOfSubject: string;
   };
   lecturers: Lecturer[];
   groups: Group[];
   startAt: Date;
   endAt: Date;
   room: {
      readonly id: string;
      room: string;
   };
}
