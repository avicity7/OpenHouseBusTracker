export type AuthedResponse = {
  Output: Array<User>,
  Tokens: {
    AccessToken: string,
    RefreshToken: string
  }
}
export type Bus = {
  BusId: string,
}

export type User = {
  Email: string,
  Role: string,
  VerificationToken: string
}


export type UserRole = {
  RoleId: number,
  Description: string 
}

export type Schedule = {
  BusScheduleId: number;
  Carplate: string;
  RouteName: string;
  DriverId: number,
  DriverName: string;
  StartTime: string;
  EndTime: string;
  Driver: Driver[];
}

export type Driver = {
	DriverId: number;
	DriverName: string;
}

export type Session = {
  Email: string,
  Role: string
} | null

// export type Schedule = {
//   BusId: number;
//   RouteId: number;
//   AssignedDriver: number;
//   StartTime: string;
//   EndTime: string;
// }
