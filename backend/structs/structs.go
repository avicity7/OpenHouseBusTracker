package structs

import "time";

type NewUser struct {
	Email    string
	Password string
	Role     int
}

type User struct {
	Email    string
	Password string
	Role     string
}

type ReturnedUser struct {
	Email string
	Role  string
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

type Schedule struct {
	BusId            int
	Carplate         string
	RouteDescription string
	DriverName       string
	StartTime        time.Time
	EndTime          time.Time
}
