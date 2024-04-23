package middleware

import (
	"fmt"
	"net/http"
	"os"
	"server/services"
	"server/structs"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func VerifyJWT(w http.ResponseWriter, r *http.Request) ([]byte, []byte) {
	secret := os.Getenv("SECRET")
	access := r.Header.Get("Access")
	err := services.VerifyAccess([]byte(access))
	if err != nil {
		refresh := r.Header.Get("Refresh")

		err = services.VerifyRefresh([]byte(refresh))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
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
				return newAccess, newRefresh
			}
		}
	}
	return nil, nil
}
