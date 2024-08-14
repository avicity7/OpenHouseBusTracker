package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"net/http"
	"server/config"
	"server/services"
	"server/structs"

	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
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
	formatted, _ := json.Marshal(output)
	w.Write(formatted)
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
	config.Cache.Delete("Users")
	config.Cache.Delete(user.Email)

	w.WriteHeader(200)
}

func UpdateSettings(w http.ResponseWriter, r *http.Request) {
	var user structs.SettingsDetails
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	err = services.UpdateSettings(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	config.Cache.Delete("Users")
	config.Cache.Delete(user.Email)

	w.WriteHeader(200)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	err := services.DeleteUser(email)

	if err != nil {
		if strings.Contains(err.Error(), "23503") {
			http.Error(w, "Cannot delete user. This user is linked to other records and cannot be deleted.", http.StatusBadRequest)
		} else {
			log.Printf("Unhandled error in controller: %v", err)
			w.WriteHeader(500)
		}
		return
	}

	config.Cache.Delete("Users")

	w.WriteHeader(200)
}
