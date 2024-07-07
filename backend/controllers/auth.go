package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/config"
	"server/middleware"
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

func BulkCreateUsers(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempFile, err := os.CreateTemp("", "users-*.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = services.BulkCreateUsers(tempFile.Name())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.Cache.Delete("Users")

	w.WriteHeader(http.StatusCreated)
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
	accessToken := r.Header.Get("Access")
	refreshToken := r.Header.Get("Refresh")
	email := chi.URLParam(r, "email")
	at, rt, err := middleware.VerifyJWT(email, accessToken, refreshToken)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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
	formatted, _ := json.Marshal(structs.AuthedResponse{Output: output, Tokens: structs.RefreshTokenResponse{AccessToken: string(at), RefreshToken: string(rt)}})
	w.Write(formatted)
}

func GetUserSettings(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	user, err := services.GetUserSettings(email)
	if err != nil {
		fmt.Printf("Error fetching user settings for email %s: %v\n", email, err)
		http.Error(w, "Failed to fetch user settings", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user settings", http.StatusInternalServerError)
		return
	}

	// config.Cache.Set(email, response, cache.DefaultExpiration) //causing issues
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	verification_token := chi.URLParam(r, "token")
	email, err := services.VerifyEmail(verification_token)
	fmt.Println(email)
	if err != nil {
		w.WriteHeader(500)
	} else {
		config.Cache.Delete(email)
		w.WriteHeader(200)
	}
}
