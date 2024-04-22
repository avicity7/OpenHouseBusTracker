package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
)

func GetJWT(w http.ResponseWriter, r *http.Request) {
	access, refresh := services.CreateJWTPair()
	result := map[string][]byte{"access": access, "refresh": refresh}

	formatted, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Failed to format into JSON.")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(formatted)
}
