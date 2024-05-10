export type AuthedResponse = {
  Output: Array<User>,
  Tokens: {
    AccessToken: string,
    RefreshToken: string
  }
}
export type Bus = {
  Carplate: string,
  Status: boolean
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

export type FollowBus = {
  ScheduleId: number,
  Carplate: string,
  DriverName: string,
  RouteName: string,
  StartTime: string,
  EndTime: string
}

export type FollowBusEvent = {
  ScheduleId: number,
  Carplate: string,
  DriverName: string,
  RouteName: string,
  Email: string,
  BusStartTime: string,
  BusEndTime: string,
  StudentStartTime: string,
  StudentEndTime: string,
}

export type EventBus = {
  Carplate: string,
  Status: boolean
} 

export type Event = {
  StopName: string,
  Order: number,
  EventId: number,
  Timestamp: string
}

export type RouteStop = {
  RouteName: string,
  StopName: string,
  Order: number
}


// export type Schedule = {
//   BusId: number;
//   RouteId: number;
//   AssignedDriver: number;
//   StartTime: string;
//   EndTime: string;
// }
