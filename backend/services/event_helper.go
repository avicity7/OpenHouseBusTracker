package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"
	// "github.com/jackc/pgx/v5"
)

func GetEventHelpers() ([]structs.EventHelper, error) {
    var eventHelpers []structs.EventHelper

    query := `
        SELECT * FROM event_helper
    `
    rows, err := config.Dbpool.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var eventHelper structs.EventHelper
        err := rows.Scan(
            &eventHelper.Carplate,
            &eventHelper.Email,
            &eventHelper.StartTime,
            &eventHelper.EndTime,
        )
        if err != nil {
            return nil, err
        }
        eventHelpers = append(eventHelpers, eventHelper)
    }
    return eventHelpers, nil
}

func CreateEventHelper(eventHelper structs.EventHelper) error {
	query := `
		INSERT INTO event_helper (carplate, email, start_time, end_time) 
		VALUES ($1, $2, $3, $4)
    `
	_, err := config.Dbpool.Exec(context.Background(), query,
		eventHelper.Carplate,
		eventHelper.Email,
		eventHelper.StartTime,
		eventHelper.EndTime,
	)

	if err != nil {
		fmt.Println("Error inserting event helper:", err)
		return err
	}

	return nil
}

// not sure why it isnt working, could be json decoding issue or struct
func UpdateEventHelper(eventHelper structs.EventHelperUpdate) error {
    query := `
        UPDATE event_helper
        SET 
            carplate = $1,
            email = $2,
            start_time = $3,
            end_time = $4
        WHERE
            carplate = $5
            AND email = $6
            AND start_time = $7
            AND end_time = $8
    `
    _, err := config.Dbpool.Exec(context.Background(), query,
        eventHelper.NewCarplate, 
        eventHelper.NewEmail,    
        eventHelper.NewStartTime,
        eventHelper.NewEndTime,  
        eventHelper.OldCarplate,
        eventHelper.OldEmail,
        eventHelper.OldStartTime,
        eventHelper.OldEndTime,
    )

    if err != nil {
        fmt.Println("Error updating event helper:", err)
        return err
    }

    return nil
}

// need to think bout how to make it take a body of ids with composite keys
func DeleteEventHelper(eventHelper structs.EventHelper) error {
    query := `
        DELETE FROM event_helper
        WHERE
            carplate = $1
            AND email = $2
            AND start_time = $3
            AND end_time = $4
    `
    _, err := config.Dbpool.Exec(context.Background(), query,
        eventHelper.Carplate,
        eventHelper.Email,
        eventHelper.StartTime,
        eventHelper.EndTime,
    )

    if err != nil {
        fmt.Println("Error deleting event helper:", err)
        return err
    }

    return nil
}