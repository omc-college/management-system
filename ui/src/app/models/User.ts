type Role = 'student' | 'lecturer' | 'parent' | 'director';
export interface User {
  id: string;
  firstName: string;
  lastName: string;
  surname: string;
  role: Role;
}
