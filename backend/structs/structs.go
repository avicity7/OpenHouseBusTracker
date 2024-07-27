package structs

import (
	"time"
)

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
	Contact           string
}

type ReturnedUser struct {
	Name              string
	Email             string
	Contact           string
	Role              string
	VerificationToken string
}

type EditUserRole struct {
	Email string
	Role  int
}

type SettingsDetails struct {
	Email   string
	Contact *string
	Name    string
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
	BusId         string
	Carplate      string
	RouteName     string
	DriverName    string
	StartTime     time.Time
	EndTime       time.Time
}

type NewSchedule struct {
	BusId     string
	RouteName string
	DriverId  int
	StartTime time.Time
	EndTime   time.Time
}

type UpdateSchedule struct {
	BusScheduleId int
	BusId         string
	Carplate      string
	RouteName     string
	DriverId      int
	StartTime     time.Time
	EndTime       time.Time
}

type UpdateScheduleRoute struct {
	BusId     string
	RouteName string
}

type ScheduleDropdownData struct {
	Buses   []EventBus
	Routes  []Route
	Drivers []Driver
}

type Driver struct {
	DriverId   int
	DriverName string
}

type FollowBus struct {
	BusId    string
	Carplate string
	Email    string
}

type EventBus struct {
	BusId    string
	Carplate string
	Status   bool
	Hidden   bool
}

type EventSchedule struct {
	ScheduleId   int
	BusId        string
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
	BusId     string
	RouteName string
	EventId   int
	StopName  string
}

type Route struct {
	RouteName string
	Color     string
}

type EventHelper struct {
	BusId    string
	Carplate string
	Name     string
	Email    string
	Shift    bool
}

type EventHelperUpdate struct {
	OldBusId string
	OldName  string
	OldShift bool

	NewBusId string
	NewName  string
	NewShift bool
}

type FollowBusEvent struct {
	ScheduleId   int
	BusId        string
	Carplate     string
	DriverName   string
	RouteName    string
	Email        string
	BusStartTime time.Time
	BusEndTime   time.Time
}

type EventHelperDropdownData struct {
	BusId    string
	Carplate string
	Name     string
}

type Stop struct {
	StopName string
	Lng      float64
	Lat      float64
}

type CurrentBus struct {
	BusId     string
	RouteName string
	Color     string
	EventType string
	Lng       float64
	Lat       float64
	Timestamp time.Time
}

type TimeDiff struct {
	DriverId       int
	StartTime      time.Time
	EndTime        time.Time
	TimeDifference time.Duration
}

type CreateMessage struct {
	From   string
	RoomId string
	Body   string
}

type Message struct {
	Timestamp time.Time
	From      string
	RoomId    string
	Body      string
}

type CreateChatRoom struct {
	User1 string
	User2 string
}

type ChatRoom struct {
	RoomId        string
	RoomName      string
	Email         string
	Name          string
	LatestMessage Message
}

type SwapRequest struct {
	From string
	With string
}

type SwapRequestResponse struct {
	From        string
	FromName    string
	With        string
	WithName    string
	TargetShift bool
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
