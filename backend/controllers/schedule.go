package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "server/services"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
    schedules, err := services.GetSchedule()
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    response, err := json.Marshal(schedules)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}
