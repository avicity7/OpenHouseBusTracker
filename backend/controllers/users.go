package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	arr, err := services.GetUsers()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
	formatted, _ := json.Marshal(arr)
	w.Write(formatted)
}

func GetUserRoles(w http.ResponseWriter, r *http.Request) {
	arr, err := services.GetUserRoles()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
	formatted, _ := json.Marshal(arr)
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

	w.WriteHeader(200)
}
