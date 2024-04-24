export type Bus = {
  BusId: string,
}

export type User = {
  Email: string,
  Role: string,
}

export type UserRole = {
  RoleId: number,
  Description: string 
}

export type Schedule = {
  BusId: number;
  Carplate: string;
  RouteName: string;
  DriverName: string;
  StartTime: string;
  EndTime: string;
}

// export type Schedule = {
//   BusId: number;
//   RouteId: number;
//   AssignedDriver: number;
//   StartTime: string;
//   EndTime: string;
// }
