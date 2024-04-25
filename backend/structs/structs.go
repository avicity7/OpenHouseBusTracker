package structs

import "time"

type NewUser struct {
	Email    string
	Password string
	Role     int
}

type User struct {
	Email             string
	Password          string
	Role              string
	VerificationToken string
}

type ReturnedUser struct {
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
	BusId      int
	Carplate   string
	RouteName  string
	DriverName string
	StartTime  time.Time
	EndTime    time.Time
}

type NewSchedule struct {
    Carplate  string    `json:"carplate"`
    RouteName string    `json:"route_name"`
    DriverId  int       `json:"driver_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}

type UpdateSchedule struct {
    BusScheduleId int       `json:"bus_schedule_id"`
    Carplate      string    `json:"carplate"`
    RouteName     string    `json:"route_name"`
    DriverId      int       `json:"driver_id"`
    StartTime     time.Time `json:"start_time"`
    EndTime       time.Time `json:"end_time"`
}

type Driver struct{
	DriverID int
	Name varchar
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