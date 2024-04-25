package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/middleware"
	"server/services"
	"server/structs"

	"github.com/patrickmn/go-cache"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Access")
	refreshToken := r.Header.Get("Refresh")
	access, refresh, err := middleware.VerifyJWT(accessToken, refreshToken)
	if err != nil {
		w.WriteHeader(500)
	} else {
		var output interface{}
		value, found := config.Cache.Get("Users")
		if found {
			output = value
		} else {
			arr, err := services.GetUsers()
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
			}
			output = arr
			config.Cache.Set("Users", arr, cache.DefaultExpiration)
		}
		w.WriteHeader(200)
		formatted, _ := json.Marshal(structs.AuthedResponse{Output: output, Tokens: structs.RefreshTokenResponse{AccessToken: string(access), RefreshToken: string(refresh)}})
		w.Write(formatted)
	}
}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	var output interface{}
	value, found := config.Cache.Get("UserRoles")
	if found {
		output = value
	} else {
		arr, err := services.GetUserRoles()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
		output = arr
		config.Cache.Set("UserRoles", arr, cache.DefaultExpiration)
	}
	w.WriteHeader(200)
	formatted, _ := json.Marshal(output)
	w.Write(formatted)
}

func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	var user structs.EditUserRole
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	err = services.UpdateUserRole(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	config.Cache.Delete(user.Email)

	w.WriteHeader(200)
}
