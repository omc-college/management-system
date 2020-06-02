type Role = 'student' | 'lecturer' | 'parent' | 'director';
export interface UserAsResource {
  id: string;
  firstName: string;
  lastName: string;
  surname: string;
  role: Role;
}
