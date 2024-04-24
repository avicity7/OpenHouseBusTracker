package middleware

import (
	"fmt"
	"os"
	"server/services"
	"server/structs"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func VerifyJWT(access string, refresh string) ([]byte, []byte, error) {
	secret := os.Getenv("SECRET")
	err := services.VerifyAccess([]byte(access))
	if err != nil {
		err = services.VerifyRefresh([]byte(refresh))
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		decode, err := jwt.Parse([]byte(refresh), jwt.WithKey(jwa.HS256, []byte(secret)))
		if err != nil {
			fmt.Println(err)
		}
		email, ok := decode.Get("Email")
		if ok {
			email_string, _ := email.(string)
			role, ok := decode.Get("Role")
			if ok {
				role_string, _ := role.(string)
				user := structs.ReturnedUser{Email: email_string, Role: role_string}
				newAccess, newRefresh := services.CreateJWTPair(user)
				return newAccess, newRefresh, nil
			}
		}
	}
	return nil, nil, nil
}
