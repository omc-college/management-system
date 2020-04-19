type Role = 'student' | 'lecturer' | 'parent' | 'director';
export interface User {
  readonly id: string;
  firstName: string;
  lastName: string;
  surname: string;
  role: Role;
}
