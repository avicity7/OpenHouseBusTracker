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
} | null

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
  StopName: string,
  Order: number
}

export type EventHelper = {
  Carplate: string,
  Email: string,
  StartTime: string,
  EndTime: string
}

// export type Schedule = {
//   BusId: number;
//   RouteId: number;
//   AssignedDriver: number;
//   StartTime: string;
//   EndTime: string;
// }
