package services

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"server/config"
	"server/structs"
	"server/utils"

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
		Claim("Name", user.Name).
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
		Claim("Name", user.Name).
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
		INSERT INTO user_table (name, email, password, role_id, verification_token) 
		VALUES (@Name, LOWER(@Email), @Password, @Role, @VerificationToken)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	verification_token := utils.GenerateRandomToken(20)

	args := pgx.NamedArgs{
		"Name":              user.Name,
		"Email":             user.Email,
		"Password":          hashedPassword,
		"Role":              user.Role,
		"VerificationToken": verification_token,
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	} else {
		utils.SendEmail(verification_token, user.Email, "", user.Name)
	}

	return nil
}

func CheckPassword(token string, password string) (bool, error) {
	var passwords []string

	query_select := `
		SELECT email FROM user_table
		WHERE verification_token = @VerificationToken
	`
	query := `
		SELECT password FROM past_passwords WHERE email = @Email
	`

	args := pgx.NamedArgs{
		"VerificationToken": token,
	}

	var username string

	err := config.Dbpool.QueryRow(context.Background(), query_select, args).Scan(&username)
	if err != nil {
		return false, err
	}

	args = pgx.NamedArgs{
		"Email": username,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var hash string
		rows.Scan(&hash)
		passwords = append(passwords, hash)
	}

	for _, hash := range passwords {
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			return true, nil
		}
	}

	return false, nil
}

func SavePassword(email string, password string) error {
	query := `
		INSERT INTO past_passwords (email, password) 
		VALUES (LOWER(@Email), @Password)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	args := pgx.NamedArgs{
		"Email":    email,
		"Password": hashedPassword,
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func BulkCreateUsers(csvFilePath string) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return err
	}

	tx, err := config.Dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(record) < 2 {
			return errors.New("invalid record length")
		}

		pwd := utils.GenerateRandomToken(6)

		user := structs.NewUser{
			Name:     record[0],
			Email:    record[1],
			Password: pwd,
			Role:     1,
		}

		err = createUserInTransaction(tx, user)
		if err != nil {
			return err
		}

		utils.SendEmail("", record[1], pwd, record[0])
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func createUserInTransaction(tx pgx.Tx, user structs.NewUser) error {
	query := `
		INSERT INTO user_table (name, email, password, role_id, verification_token) 
		VALUES ($1, $2, $3, $4, $5)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	verificationToken := ""

	_, err = tx.Exec(context.Background(), query, user.Name, user.Email, hashedPassword, user.Role, verificationToken)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email string) (structs.ReturnedUser, error) {
	var user structs.ReturnedUser
	query := `
		SELECT name, email, role_name, verification_token, COALESCE(contact, '') as contact FROM user_table 
		JOIN user_role ON user_table.role_id = user_role.role_id 
		WHERE email = @Email
	`
	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Name, &user.Email, &user.Role, &user.VerificationToken, &user.Contact)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	return user, nil
}

func GetUserSettings(email string) (structs.SettingsDetails, error) {
	var user structs.SettingsDetails
	query := `
		SELECT name, contact, email FROM user_table 
		WHERE email = @Email
	`
	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Name, &user.Contact, &user.Email)
	if err != nil {
		return structs.SettingsDetails{}, err
	}

	return user, nil
}

func Login(login structs.Login) (structs.ReturnedUser, error) {
	var user structs.User
	query := `
		SELECT name, email, password, role_name, verification_token FROM user_table 
		JOIN user_role ON user_table.role_id = user_role.role_id 
		WHERE email = LOWER(@Email)
	`
	args := pgx.NamedArgs{
		"Email": login.Email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Name, &user.Email, &user.Password, &user.Role, &user.VerificationToken)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	returnedUser := structs.ReturnedUser{
		Name:              user.Name,
		Email:             user.Email,
		Role:              user.Role,
		VerificationToken: user.VerificationToken,
	}

	return returnedUser, nil
}

func VerifyEmail(verification_token string) (string, error) {
	query_select := `
		SELECT email FROM user_table
		WHERE verification_token = @VerificationToken
	`
	query_update := `
		UPDATE user_table
		SET verification_token = ''
		WHERE verification_token = @VerificationToken
	`
	args := pgx.NamedArgs{
		"VerificationToken": verification_token,
	}
	var username string
	err := config.Dbpool.QueryRow(context.Background(), query_select, args).Scan(&username)
	if err != nil {
		return "", err
	}
	_, err = config.Dbpool.Exec(context.Background(), query_update, args)
	if err != nil {
		return "", err
	}
	return username, nil
}

func ResetPassword(token string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	query_select := `
		SELECT email FROM user_table
		WHERE verification_token = @VerificationToken
	`
	query_update := `
		UPDATE user_table
		SET verification_token = ''
		WHERE verification_token = @VerificationToken
	`
	query_password := `
		UPDATE user_table
		SET password = @Password
		WHERE verification_token = @VerificationToken
	`
	args := pgx.NamedArgs{
		"VerificationToken": token,
		"Password":          hashedPassword,
	}

	var username string

	err = config.Dbpool.QueryRow(context.Background(), query_select, args).Scan(&username)
	if err != nil {
		return err
	}

	_, err = config.Dbpool.Exec(context.Background(), query_password, args)
	if err != nil {
		return err
	}

	_, err = config.Dbpool.Exec(context.Background(), query_update, args)
	if err != nil {
		return err
	}

	SavePassword(username, password)

	return nil
}

func StartResetPassword(email string) error {
	user, err := GetUser(email)
	if err != nil {
		return err
	}

	token := "reset" + utils.GenerateRandomToken(20)

	query_update := `
		UPDATE user_table
		SET verification_token = @VerificationToken
		WHERE email = @Email
	`
	args := pgx.NamedArgs{
		"VerificationToken": token,
		"Email":             email,
	}

	_, err = config.Dbpool.Exec(context.Background(), query_update, args)
	if err != nil {
		return err
	}

	utils.SendEmail(token, email, "reset", user.Name)

	return nil
}
