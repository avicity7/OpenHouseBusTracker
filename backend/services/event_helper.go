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
        SELECT carplate, ut.name, shift FROM event_helper eh
		JOIN user_table ut ON ut.email = eh.email
		ORDER BY carplate ASC, ut.name ASC
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
			&eventHelper.Name,
			&eventHelper.Shift,
		)
		if err != nil {
			return nil, err
		}
		eventHelpers = append(eventHelpers, eventHelper)
	}
	return eventHelpers, nil
}

func CreateEventHelpers(eventHelpers []structs.EventHelper) error {
    tx, err := config.Dbpool.Begin(context.Background())
    if err != nil {
        fmt.Println("Failed to begin transaction:", err)
        return err
    }
    defer func() {
        if err != nil {
            tx.Rollback(context.Background())
        } else {
            tx.Commit(context.Background())
        }
    }()

    for _, eventHelper := range eventHelpers {
        email, err := GetEmail(eventHelper.Name)
        if err != nil {
            fmt.Println("Error retrieving email:", err)
            return err
        }

        query := `
            INSERT INTO event_helper (carplate, email, shift) 
            VALUES ($1, $2, $3)
        `
        _, err = tx.Exec(context.Background(), query,
            eventHelper.Carplate,
            email,
            eventHelper.Shift,
        )
        if err != nil {
            fmt.Println("Error inserting event helper:", err)
            return err
        }
        fmt.Printf("Event helper inserted successfully for email: %s\n", email)
    }

    return nil
}

func EventHelperExists(email string) (bool, error) {
    var exists bool
    query := `
        SELECT EXISTS (
            SELECT 1 FROM event_helper
            WHERE email = $1
        )
    `
    err := config.Dbpool.QueryRow(context.Background(), query, email).Scan(&exists)
    if err != nil {
        fmt.Println("Error checking existing event helper:", err)
        return false, err
    }

    if exists {
        fmt.Printf("Event helper already exists for email: %s\n", email)
    } else {
        fmt.Printf("No existing event helper found for email: %s\n", email)
    }

    return exists, nil
}

func UpdateEventHelper(eventHelper structs.EventHelperUpdate) error {
	newEmail, err := GetEmail(eventHelper.NewName)
	if err != nil {
		return err
	}

	oldEmail, err := GetEmail(eventHelper.OldName)
	if err != nil {
		return err
	}
	query := `
        UPDATE event_helper
        SET 
            carplate = $1,
            email = $2,
            shift = $3
        WHERE
            carplate = $4
            AND email = $5
            AND shift = $6
    `
	_, err = config.Dbpool.Exec(context.Background(), query,
		eventHelper.NewCarplate,
		newEmail,
		eventHelper.NewShift,
		eventHelper.OldCarplate,
		oldEmail,
		eventHelper.OldShift,
	)

	fmt.Println("new email ", newEmail)
	fmt.Println("old email ", oldEmail)

	if err != nil {
		fmt.Println("Error updating event helper:", err)
		return err
	}

	return nil
}

// need to think bout how to make it take a body of ids with composite keys
func DeleteEventHelper(eventHelper structs.EventHelper) error {
	email, err := GetEmail(eventHelper.Name)
	if err != nil {
		return err
	}

	query := `
        DELETE FROM event_helper
        WHERE
            carplate = $1
            AND email = $2
            AND shift = $3
    `
	_, err = config.Dbpool.Exec(context.Background(), query,
		eventHelper.Carplate,
		email,
		eventHelper.Shift,
	)
	if err != nil {
		fmt.Println("Error deleting event helper:", err)
		return err
	}

	return nil
}

func GetEventHelperDropdownData() ([]structs.EventHelperDropdownData, error) {
	var dropdownData []structs.EventHelperDropdownData

	query := `
        SELECT 
            b.carplate,
            NULL AS name
        FROM 
            bus b

        UNION ALL

        SELECT 
            NULL AS carplate,
            u.name
        FROM 
            user_table u
        WHERE 
            u.role_id = 1
            AND u.name NOT IN (
                SELECT 
                    ut.name 
                FROM 
                    event_helper eh
                JOIN 
                    user_table ut ON ut.email = eh.email
            )
		ORDER BY 
			carplate ASC, name ASC
    `

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data structs.EventHelperDropdownData
		err := rows.Scan(
			&data.Carplate,
			&data.Name,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		dropdownData = append(dropdownData, data)

	}

	return dropdownData, nil
}
