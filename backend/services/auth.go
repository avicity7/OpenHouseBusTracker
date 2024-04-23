package services

import (
	"context"
	"fmt"
	"os"
	"server/config"
	"server/structs"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccessToken(user structs.ReturnedUser) []byte {
	secret := os.Getenv("SECRET")

	access_token, err := jwt.NewBuilder().
		Issuer("server").
		Claim("Email", user.Email).
		Claim("Role", user.Role).
		Expiration(time.Now().Add(24 * time.Second)).
		Build()
	if err != nil {
		fmt.Println("Access Token generation failed.")
	}

	signed_access_token, err := jwt.Sign(access_token, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		fmt.Println(err)
	}

	return signed_access_token
}

func CreateRefreshToken(user structs.ReturnedUser) []byte {
	secret := os.Getenv("SECRET")

	refresh_token, err := jwt.NewBuilder().
		Issuer("server").
		Claim("Email", user.Email).
		Claim("Role", user.Role).
		Expiration(time.Now().Add(168 * time.Hour)).
		Build()
	if err != nil {
		fmt.Println("Access Token generation failed.")
	}

	signed_refresh_token, err := jwt.Sign(refresh_token, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		fmt.Println(err)
	}

	return signed_refresh_token
}

func CreateJWTPair(user structs.ReturnedUser) ([]byte, []byte) {
	access := CreateAccessToken(user)
	refresh := CreateRefreshToken(user)

	return access, refresh
}

func VerifyRefresh(refresh []byte) error {
	secret := os.Getenv("SECRET")

	parsed_refresh, err := jwt.Parse(refresh, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		return err
	}
	_ = parsed_refresh
	return nil
}

func VerifyAccess(access []byte) error {
	secret := os.Getenv("SECRET")

	parsed_access, err := jwt.Parse(access, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		return err
	}
	_ = parsed_access
	return nil
}

func CreateUser(user structs.NewUser) error {
	query := `
		INSERT INTO "User" ("Email", "Password", "Role") 
		VALUES (@Email, @Password, @Role)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	args := pgx.NamedArgs{
		"Email":    user.Email,
		"Password": hashedPassword,
		"Role":     user.Role,
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email string) (structs.ReturnedUser, error) {
	var user structs.ReturnedUser
	query := `
		SELECT "Email", "Description" as "Role" FROM "User" 
		JOIN "UserRole" ON "User"."Role" = "UserRole"."RoleId" 
		WHERE "Email" = @Email
	`
	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Email, &user.Role)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	return user, nil
}

func Login(login structs.Login) (structs.ReturnedUser, error) {
	var user structs.User
	query := `
		SELECT "Email", "Password", "Description" as "Role" FROM "User" 
		JOIN "UserRole" ON "User"."Role" = "UserRole"."RoleId" 
		WHERE "Email" = @Email
	`
	args := pgx.NamedArgs{
		"Email": login.Email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Email, &user.Password, &user.Role)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	returnedUser := structs.ReturnedUser{
		Email: user.Email,
		Role:  user.Role,
	}

	return returnedUser, nil
}
