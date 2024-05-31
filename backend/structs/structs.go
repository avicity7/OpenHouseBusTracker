package structs

import "time"

type NewUser struct {
	Name     string
	Email    string
	Password string
	Role     int
}

type User struct {
	Name              string
	Email             string
	Password          string
	Role              string
	VerificationToken string
}

type ReturnedUser struct {
	Name              string
	Email             string
	Role              string
	VerificationToken string
}

type EditUserRole struct {
	Email string
	Role  int
}

type ReturnedUserArray []ReturnedUser

type UserRole struct {
	RoleId      int
	Description string
}

type UserRoleArray []UserRole

type Login struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User         ReturnedUser
	AccessToken  string
	RefreshToken string
}

type RefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthedResponse struct {
	Output interface{}
	Tokens RefreshTokenResponse
}

type Schedule struct {
	BusScheduleId int
	Carplate      string
	RouteName     string
	DriverName    string
	StartTime     time.Time
	EndTime       time.Time
}

type NewSchedule struct {
	Carplate  string
	RouteName string
	DriverId  int
	StartTime time.Time
	EndTime   time.Time
}

type UpdateSchedule struct {
	BusScheduleId int
	Carplate      string
	RouteName     string
	DriverId      int
	StartTime     time.Time
	EndTime       time.Time
}

type ScheduleDropdownData struct {
	Carplate  string
	RouteName string
	Driver    []Driver
}

type Driver struct {
	DriverId   int
	DriverName string
}

type FollowBus struct {
	Carplate string
	Email    string
}

type EventBus struct {
	Carplate string
	Status   bool
}

type EventSchedule struct {
	ScheduleId   int
	Carplate     string
	DriverName   string
	RouteName    string
	BusStartTime time.Time
	BusEndTime   time.Time
}

type RouteStep struct {
	RouteName string
	StopName  string
	Order     int
	Lng       float64
	Lat       float64
}

type Event struct {
	StopName  string
	Order     int
	EventId   int
	Timestamp time.Time
}

type EventInput struct {
	Carplate  string
	RouteName string
	EventId   int
	StopName  string
}

type Route struct {
	RouteName string
	Color     string
}

type EventHelper struct {
	Carplate string
	Name     string
	Shift    bool
}

type EventHelperUpdate struct {
	OldCarplate string
	OldName     string
	OldShift    bool

	NewCarplate string
	NewName     string
	NewShift    bool
}

type FollowBusEvent struct {
	ScheduleId       int
	Carplate         string
	DriverName       string
	RouteName        string
	Email            string
	BusStartTime     time.Time
	BusEndTime       time.Time
	StudentStartTime time.Time
	StudentEndTime   time.Time
}

type EventHelperDropdownData struct {
	Carplate *string
	Name     *string
}

type Stop struct {
	StopName string
	Lng      float64
	Lat      float64
}

type CurrentBus struct {
	Carplate  string
	RouteName string
	Color     string
	EventType string
	Lng       float64
	Lat       float64
	Timestamp time.Time
}

// could try to use clearer seperation of structs / reduce structs needed

// type Schedule struct {
// 	BusId      int
// 	Carplate   string
// 	RouteName  string
// 	DriverName string
// 	StartTime  time.Time
// 	EndTime    time.Time
// }

// type ScheduleRequest struct {
// 	Carplate  string    `json:"carplate"`
// 	RouteName string    `json:"route_name"`
// 	DriverId  int       `json:"driver_id"`
// 	StartTime time.Time `json:"start_time"`
// 	EndTime   time.Time `json:"end_time"`
// }

// type ScheduleUpdateRequest struct {
// 	BusId     int       `json:"bus_id"`
// 	Carplate  string    `json:"carplate"`
// 	RouteName string    `json:"route_name"`
// 	DriverId  int       `json:"driver_id"`
// 	StartTime time.Time `json:"start_time"`
// 	EndTime   time.Time `json:"end_time"`
// }

//ref
//type Route struct {
//    RouteName string `json:"route_name"`
//}
