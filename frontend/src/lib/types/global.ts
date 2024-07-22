export type AuthedResponse = {
  Output: Array<User>,
  Tokens: {
    AccessToken: string,
    RefreshToken: string
  }
}

export type User = {
  Name: string,
  Email: string,
  Role: string,
  VerificationToken: string,
  Contact: string
}

export type UserRole = {
  RoleId: number,
  Description: string 
}

export type Schedule = {
  BusScheduleId: number,
  BusId: string,
  Carplate: string,
  RouteName: string,
  DriverId: number,
  DriverName: string,
  StartTime: string,
  EndTime: string,
  Driver: Driver[]
}

export type DropdownData = {
  Buses:  EventBus[],
  Routes: Route[],
  Drivers: Driver[]
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
  BusId: string,
  Carplate: string,
  DriverName: string,
  RouteName: string,
  StartTime: string,
  EndTime: string
}

export type FollowBusEvent = {
  ScheduleId: number,
  BusId: string,
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
  BusId: string,
  Carplate: string,
  Status: boolean,
  Hidden: boolean
} 

export type Event = {
  StopName: string,
  Order: number,
  EventId: number,
  Timestamp: string
}

export type RouteStep = {
  RouteName: string,
  StopName: string,
  Order: number,
  Lng: number,
  Lat: number
}

export type Route = {
  RouteName: string,
  Color: string
}

export type EventHelper = {
  BusId: string,
  Carplate: string,
  Name: string,
  Shift: boolean
}

export type Stop = {
  StopName: string,
  Lng: number,
  Lat: number,
}

export type CurrentBus = {
  BusId: string,
  Carplate: string,
  RouteName: string,
  Color: string,
  EventType: string,
  Lat: number,
  Lng: number,
  Timestamp: string
}

export type LoginResponse = {
  User: User,
  AccessToken: string, 
  RefreshToken: string
}

export type Message = {
  Timestamp: string, 
  From: string, 
  RoomId: string, 
  Body: string
}

export type ChatRoom = {
  RoomId: string,
  User1: string,
  User2: string
  LatestMessage: Message
}

export type RouteStates = {
  [key: string]: string;
};

export type Demand = {
  [key: string]: number;
};

export type BusAssignments = {
  [key: string]: string[];
};

// export type Schedule = {
//   BusId: number;
//   RouteId: number;
//   AssignedDriver: number;
//   StartTime: string;
//   EndTime: string;
// }
