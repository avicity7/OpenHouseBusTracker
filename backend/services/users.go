package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func GetEmail(name string) (string, error) {
	var email string
	query := `
		SELECT email FROM user_table
		WHERE name = @Name
	`

	args := pgx.NamedArgs{
		"Name": name,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&email)
	if err != nil {
		return "", err
	}

	return email, nil
}

func GetUsers() (structs.ReturnedUserArray, error) {
	output := make(structs.ReturnedUserArray, 0)
	query := `
		SELECT name, email, COALESCE(contact, '') as contact, role_name, verification_token FROM user_table 
		JOIN user_role ON user_table.role_id = user_role.role_id 
		ORDER BY email ASC
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return structs.ReturnedUserArray{}, err
	}

	for rows.Next() {
		var user structs.ReturnedUser
		rows.Scan(&user.Name, &user.Email, &user.Contact, &user.Role, &user.VerificationToken)
		output = append(output, user)
	}

	return output, nil
}

func GetUserRoles() (structs.UserRoleArray, error) {
	output := make(structs.UserRoleArray, 0)
	query := `
		SELECT * FROM user_role
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
		UPDATE user_table
		SET role_id = @Role
		WHERE email = @Email
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

func UpdateSettings(user structs.SettingsDetails) error {
	query := `
		UPDATE user_table
		SET name = @Name, contact = @Contact
		WHERE email = @Email
	`

	args := pgx.NamedArgs{
		"Name":          user.Name,
		"Contact": 		 user.Contact,
		"Email":         user.Email,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(email string) error {
	query := `
		DELETE FROM user_table ut
		WHERE email = @Email
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}
