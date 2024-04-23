package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func GetUsers() (structs.ReturnedUserArray, error) {
	output := make(structs.ReturnedUserArray, 0)
	query := `
		SELECT "Email", "Description" as "Role" FROM "User" 
		JOIN "UserRole" ON "User"."Role" = "UserRole"."RoleId" 
		ORDER BY "Email" ASC
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return structs.ReturnedUserArray{}, err
	}

	for rows.Next() {
		var user structs.ReturnedUser
		rows.Scan(&user.Email, &user.Role)
		output = append(output, user)
	}

	return output, nil
}

func GetUserRoles() (structs.UserRoleArray, error) {
	output := make(structs.UserRoleArray, 0)
	query := `
		SELECT * FROM "UserRole"
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return structs.UserRoleArray{}, err
	}

	for rows.Next() {
		var userRole structs.UserRole
		rows.Scan(&userRole.RoleId, &userRole.Description)
		output = append(output, userRole)
	}

	return output, nil
}

func UpdateUserRole(user structs.EditUserRole) error {
	query := `
		UPDATE "User"
		SET "Role" = @Role
		WHERE "Email" = @Email
	`

	args := pgx.NamedArgs{
		"Role":  user.Role,
		"Email": user.Email,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}
