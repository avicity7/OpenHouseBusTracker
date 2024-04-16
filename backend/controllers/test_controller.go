package controllers

import (
	"net/http"
	"server/services"
)

func GetTestController(w http.ResponseWriter, r *http.Request) {
	result := services.GetTestService()
	w.Header().Set("Content-Type", "application/json")
	// render.JSON(w, r, result)
	w.Write(result)
}
