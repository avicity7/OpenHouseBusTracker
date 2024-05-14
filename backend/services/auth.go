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
	// "strconv"
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
		VALUES (@Name, @Email, @Password, @Role, @VerificationToken)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	verification_token := utils.GenerateRandomToken(20)

	args := pgx.NamedArgs{
		"Name": 			 user.Name,
		"Email":             user.Email,
		"Password":          hashedPassword,
		"Role":              user.Role,
		"VerificationToken": verification_token,
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	} else {
		utils.SendEmail(verification_token, user.Email)
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

		if len(record) < 3 {
			return errors.New("invalid record length")
		}

		// role, err := strconv.Atoi(record[3])
		// if err != nil {
		// 	return err
		// }

		user := structs.NewUser{
			Name:     record[0],
			Email:    record[1],
			Password: record[2],
			Role:     0,
		}

		err = createUserInTransaction(tx, user)
		if err != nil {
			return err
		}
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

	verificationToken := utils.GenerateRandomToken(20)

	_, err = tx.Exec(context.Background(), query, user.Name, user.Email, hashedPassword, user.Role, verificationToken)
	if err != nil {
		return err
	}

	utils.SendEmail(verificationToken, user.Email)
	return nil
}

func GetUser(email string) (structs.ReturnedUser, error) {
	var user structs.ReturnedUser
	query := `
		SELECT name, email, role_name, verification_token FROM user_table 
		JOIN user_role ON user_table.role_id = user_role.role_id 
		WHERE email = @Email
	`
	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Name, &user.Email, &user.Role, &user.VerificationToken)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	return user, nil
}

func Login(login structs.Login) (structs.ReturnedUser, error) {
	var user structs.User
	query := `
		SELECT name, email, password, role_name, verification_token FROM user_table 
		JOIN user_role ON user_table.role_id = user_role.role_id 
		WHERE email = @Email
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
		Name: 			   user.Name,	
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
	fmt.Println(username)
	_, err = config.Dbpool.Exec(context.Background(), query_update, args)
	if err != nil {
		return "", err
	}
	return username, nil
}


