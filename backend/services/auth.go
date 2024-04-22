package services

import (
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func CreateAccessToken() []byte {
	secret := os.Getenv("SECRET")

	access_token, err := jwt.NewBuilder().
		Issuer("me").
		Claim("foo", "bar").
		Expiration(time.Now().Add(24 * time.Hour)).
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

func CreateRefreshToken() []byte {
	secret := os.Getenv("SECRET")

	refresh_token, err := jwt.NewBuilder().
		Issuer("me").
		Claim("foo", "bar").
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

func CreateJWTPair() ([]byte, []byte) {
	access := CreateAccessToken()
	refresh := CreateRefreshToken()

	return access, refresh
}

func VerifyJWTPair(access []byte, refresh []byte) {
	secret := os.Getenv("SECRET")

	parsed_access, err := jwt.Parse(access, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		fmt.Printf("failed to parse JWT: %s\n", err)
	}
	_ = parsed_access

	parsed_refresh, err := jwt.Parse(refresh, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		fmt.Printf("failed to parse JWT: %s\n", err)
	}
	_ = parsed_refresh
}
