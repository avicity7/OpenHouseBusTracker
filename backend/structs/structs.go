// Package structs defines the structs used throughout the API.
package structs

import (
	"time"
)

// NewUser represents the structure of a new user.
type NewUser struct {
	Name     string // Name of the user.
	Email    string // Email of the user.
	Password string // Password of the user.
	Role     int    // Role of the user.
}

// User represents the structure of a user.
type User struct {
	Name              string // Name of the user.
	Email             string // Email of the user.
	Password          string // Password of the user.
	Role              string // Role of the user.
	VerificationToken string // Verification token of the user.
	Contact           string // Contact information of the user.
}

// ReturnedUser represents the structure of a returned user.
type ReturnedUser struct {
	Name              string // Name of the user.
	Email             string // Email of the user.
	Contact           string // Contact information of the user.
	Role              string // Role of the user.
	VerificationToken string // Verification token of the user.
}

// EditUserRole represents the structure for editing user role.
type EditUserRole struct {
	Email string // Email of the user.
	Role  int    // Role of the user.
}

// SettingsDetails represents the structure for user settings details.
type SettingsDetails struct {
	Email   string  // Email of the user.
	Contact *string // Contact information of the user.
	Name    string  // Name of the user.
}

// ReturnedUserArray represents an array of returned users.
type ReturnedUserArray []ReturnedUser

// UserRole represents the structure of a user role.
type UserRole struct {
	RoleId      int    // ID of the role.
	Description string // Description of the role.
}

// UserRoleArray represents an array of user roles.
type UserRoleArray []UserRole

// Login represents the structure for user login.
type Login struct {
	Email    string // Email of the user.
	Password string // Password of the user.
}

// LoginResponse represents the structure for login response.
type LoginResponse struct {
	User         ReturnedUser // Returned user information.
	AccessToken  string       // Access token for authentication.
	RefreshToken string       // Refresh token for authentication.
}

// RefreshTokenResponse represents the structure for refresh token response.
type RefreshTokenResponse struct {
	AccessToken  string // Access token for authentication.
	RefreshToken string // Refresh token for authentication.
}

// AuthedResponse represents the structure for authenticated response.
type AuthedResponse struct {
	Output interface{}          // Output data.
	Tokens RefreshTokenResponse // Tokens for authentication.
}

// Schedule represents the structure of a bus schedule.
type Schedule struct {
	BusScheduleId int       // ID of the bus schedule.
	BusId         string    // ID of the bus.
	Carplate      string    // Carplate of the bus.
	RouteName     string    // Name of the route.
	DriverName    string    // Name of the driver.
	StartTime     time.Time // Start time of the schedule.
	EndTime       time.Time // End time of the schedule.
}

// NewSchedule represents the structure for creating a new bus schedule.
type NewSchedule struct {
	BusId     string    // ID of the bus.
	RouteName string    // Name of the route.
	DriverId  int       // ID of the driver.
	StartTime time.Time // Start time of the schedule.
	EndTime   time.Time // End time of the schedule.
}

// UpdateSchedule represents the structure for updating a bus schedule.
type UpdateSchedule struct {
	BusScheduleId int       // ID of the bus schedule.
	BusId         string    // ID of the bus.
	Carplate      string    // Carplate of the bus.
	RouteName     string    // Name of the route.
	DriverId      int       // ID of the driver.
	StartTime     time.Time // Start time of the schedule.
	EndTime       time.Time // End time of the schedule.
}

// UpdateScheduleRoute represents the structure for updating a bus schedule route.
type UpdateScheduleRoute struct {
	BusId     string // ID of the bus.
	RouteName string // Name of the route.
}

// ScheduleDropdownData represents the structure for dropdown data in the schedule.
type ScheduleDropdownData struct {
	Buses   []EventBus // List of event buses.
	Routes  []Route    // List of routes.
	Drivers []Driver   // List of drivers.
}

// Driver represents the structure of a driver.
type Driver struct {
	DriverId   int    // ID of the driver.
	DriverName string // Name of the driver.
}

// FollowBus represents the structure for following a bus.
type FollowBus struct {
	BusId    string // ID of the bus.
	Carplate string // Carplate of the bus.
	Email    string // Email of the user.
}

// EventBus represents the structure of an event bus.
type EventBus struct {
	BusId    string // ID of the bus.
	Carplate string // Carplate of the bus.
	Status   bool   // Status of the bus.
	Hidden   bool   // Hidden status of the bus.
}

// EventSchedule represents the structure of an event schedule.
type EventSchedule struct {
	ScheduleId   int       // ID of the schedule.
	BusId        string    // ID of the bus.
	Carplate     string    // Carplate of the bus.
	DriverName   string    // Name of the driver.
	RouteName    string    // Name of the route.
	BusStartTime time.Time // Start time of the bus.
	BusEndTime   time.Time // End time of the bus.
}

// RouteStep represents the structure of a route step.
type RouteStep struct {
	RouteName string  // Name of the route.
	StopName  string  // Name of the stop.
	Order     int     // Order of the step.
	Lng       float64 // Longitude of the step.
	Lat       float64 // Latitude of the step.
}

// Event represents the structure of an event.
type Event struct {
	StopName  string    // Name of the stop.
	Order     int       // Order of the event.
	EventId   int       // ID of the event.
	Timestamp time.Time // Timestamp of the event.
}

// EventInput represents the structure for event input.
type EventInput struct {
	BusId     string // ID of the bus.
	RouteName string // Name of the route.
	EventId   int    // ID of the event.
	StopName  string // Name of the stop.
}

// Route represents the structure of a route.
type Route struct {
	RouteName string // Name of the route.
	Color     string // Color of the route.
}

// EventHelper represents the structure of an event helper.
type EventHelper struct {
	BusId    string // ID of the bus.
	Carplate string // Carplate of the bus.
	Name     string // Name of the helper.
	Email    string // Email of the helper.
	Shift    bool   // Shift status of the helper.
}

// EventHelperUpdate represents the structure for updating an event helper.
type EventHelperUpdate struct {
	OldBusId string // Old ID of the bus.
	OldName  string // Old name of the helper.
	OldShift bool   // Old shift status of the helper.

	NewBusId string // New ID of the bus.
	NewName  string // New name of the helper.
	NewShift bool   // New shift status of the helper.
}

// FollowBusEvent represents the structure for following a bus event.
type FollowBusEvent struct {
	ScheduleId   int       // ID of the schedule.
	BusId        string    // ID of the bus.
	Carplate     string    // Carplate of the bus.
	DriverName   string    // Name of the driver.
	RouteName    string    // Name of the route.
	Email        string    // Email of the user.
	BusStartTime time.Time // Start time of the bus.
	BusEndTime   time.Time // End time of the bus.
}

// EventHelperDropdownData represents the structure for dropdown data in the event helper.
type EventHelperDropdownData struct {
	BusId    string // ID of the bus.
	Carplate string // Carplate of the bus.
	Name     string // Name of the helper.
}

// Stop represents the structure of a stop.
type Stop struct {
	StopName string  // Name of the stop.
	Lng      float64 // Longitude of the stop.
	Lat      float64 // Latitude of the stop.
}

// CurrentBus represents the structure of a current bus.
type CurrentBus struct {
	BusId     string    // ID of the bus.
	Carplate  string    // Carplate of the bus.
	RouteName string    // Name of the route.
	Color     string    // Color of the route.
	EventType string    // Type of the event.
	Lng       float64   // Longitude of the bus.
	Lat       float64   // Latitude of the bus.
	Timestamp time.Time // Timestamp of the bus.
}

// TimeDiff represents the structure for time difference.
type TimeDiff struct {
	DriverId       int           // ID of the driver.
	StartTime      time.Time     // Start time of the schedule.
	EndTime        time.Time     // End time of the schedule.
	TimeDifference time.Duration // Time difference between start and end time.
}

// CreateMessage represents the structure for creating a message.
type CreateMessage struct {
	From   string // Sender of the message.
	RoomId string // ID of the chat room.
	Body   string // Body of the message.
}

// Message represents the structure of a message.
type Message struct {
	Timestamp time.Time // Timestamp of the message.
	From      string    // Sender of the message.
	FromName  string    // Name of the sender.
	RoomId    string    // ID of the chat room.
	Body      string    // Body of the message.
}

// CreateChatRoom represents the structure for creating a chat room.
type CreateChatRoom struct {
	User1 string // User 1 in the chat room.
	User2 string // User 2 in the chat room.
}

// ChatRoom represents the structure of a chat room.
type ChatRoom struct {
	RoomId        string  // ID of the chat room.
	RoomName      string  // Name of the chat room.
	Email         string  // Email of the user.
	Name          string  // Name of the user.
	LatestMessage Message // Latest message in the chat room.
}

// SwapRequest represents the structure for a swap request.
type SwapRequest struct {
	From string // Sender of the swap request.
	With string // Receiver of the swap request.
}

// SwapRequestResponse represents the structure for a swap request response.
type SwapRequestResponse struct {
	Timestamp   time.Time // Timestamp of the swap request.
	From        string    // Sender of the swap request.
	FromName    string    // Name of the sender.
	With        string    // Receiver of the swap request.
	WithName    string    // Name of the receiver.
	TargetShift bool      // Target shift status.
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
