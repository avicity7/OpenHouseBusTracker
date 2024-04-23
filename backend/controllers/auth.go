package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/services"
	"server/structs"

	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser structs.NewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = services.CreateUser(newUser)
	if err != nil {
		fmt.Println(err)
	}
	login := structs.Login{
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	user, err := services.Login(login)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	access, refresh := services.CreateJWTPair(user)
	formatted, _ := json.Marshal(structs.LoginResponse{User: user, AccessToken: string(access), RefreshToken: string(refresh)})
	config.Cache.Delete("Users")
	w.WriteHeader(200)
	w.Write(formatted)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login structs.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(500)
	}
	user, err := services.Login(login)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	} else {
		access, refresh := services.CreateJWTPair(user)
		formatted, _ := json.Marshal(structs.LoginResponse{User: user, AccessToken: string(access), RefreshToken: string(refresh)})
		w.WriteHeader(200)
		w.Write(formatted)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var output interface{}
	email := chi.URLParam(r, "email")
	value, found := config.Cache.Get(email)
	if found {
		output = value
	} else {
		user, err := services.GetUser(email)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
		output = user
		config.Cache.Set(email, user, cache.DefaultExpiration)
	}
	formatted, _ := json.Marshal(output)
	w.Write(formatted)
}

func VerifyJWT(w http.ResponseWriter, r *http.Request) {
	access, err := r.Cookie("accessToken")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	var user structs.ReturnedUser
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	err = services.VerifyAccess([]byte(access.Value))
	if err != nil {
		refresh, err := r.Cookie("refreshToken")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}

		err = services.VerifyRefresh([]byte(refresh.Value))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
		newAccess, newRefresh := services.CreateJWTPair(user)
		formatted, _ := json.Marshal(structs.RefreshTokenResponse{AccessToken: string(newAccess), RefreshToken: string(newRefresh)})
		w.WriteHeader(200)
		w.Write(formatted)
	}
}
