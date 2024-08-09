package services

import (
	"context"
	"errors"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5"
)

func GetEventHelpers() ([]structs.EventHelper, error) {
	var eventHelpers []structs.EventHelper

	query := `
		SELECT eh.bus_id, carplate, ut.name, ut.email, shift FROM event_helper eh
		JOIN bus b ON eh.bus_id = b.bus_id
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
			&eventHelper.BusId,
			&eventHelper.Carplate,
			&eventHelper.Name,
			&eventHelper.Email,
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
            INSERT INTO event_helper (bus_id, email, shift) 
            VALUES ($1, $2, $3)
        `
		_, err = tx.Exec(context.Background(), query,
			eventHelper.BusId,
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

func BulkCreateEventHelpers(eventHelpers []structs.EventHelper) error {
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
		query := `
            INSERT INTO event_helper (bus_id, email, shift) 
            VALUES ($1, $2, $3)
        `
		_, err = tx.Exec(context.Background(), query,
			eventHelper.BusId,
			eventHelper.Email,
			eventHelper.Shift,
		)
		if err != nil {
			fmt.Println("Error inserting event helper:", err)
			return err
		}
		fmt.Printf("Event helper inserted successfully for email: %s\n", eventHelper.Email)
	}

	return nil
}

func GetBusIdByCarplate(carplate string) (string, error) {
	var busId string

	query := `SELECT bus_id FROM bus WHERE carplate = $1`
	err := config.Dbpool.QueryRow(context.Background(), query, carplate).Scan(&busId)
	if err != nil {
		return "", errors.New("bus ID not found")
	}

	return busId, nil
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
            bus_id = $1,
            email = $2,
            shift = $3
        WHERE
            bus_id = $4
            AND email = $5
            AND shift = $6
    `
	_, err = config.Dbpool.Exec(context.Background(), query,
		eventHelper.NewBusId,
		newEmail,
		eventHelper.NewShift,
		eventHelper.OldBusId,
		oldEmail,
		eventHelper.OldShift,
	)

	if err != nil {
		fmt.Println("Error updating event helper:", err)
		return err
	}

	return nil
}

func DeleteEventHelper(eventHelper structs.EventHelper) error {
	email, err := GetEmail(eventHelper.Name)
	if err != nil {
		return err
	}

	query := `
        DELETE FROM event_helper
        WHERE
            bus_id = $1
            AND email = $2
            AND shift = $3
    `
	_, err = config.Dbpool.Exec(context.Background(), query,
		eventHelper.BusId,
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
        b.bus_id,
        b.carplate,
        '' AS name
    FROM 
        bus b

    UNION ALL

    SELECT
        '00000000-0000-0000-0000-000000000000'::UUID AS bus_id,
        '' AS carplate,
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
			&data.BusId,
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

func GetAvailableSwaps(email string) ([]structs.EventHelper, error) {
	var eventHelpers []structs.EventHelper

	query := `
		WITH sq AS (
    		SELECT email, shift FROM event_helper eh WHERE eh.email = @Email
		)
		SELECT eh.bus_id, carplate, ut.name, ut.email, eh.shift FROM event_helper eh 
    	JOIN bus b ON eh.bus_id = b.bus_id
		JOIN user_table ut ON ut.email = eh.email
		RIGHT JOIN sq ON NOT eh.shift = sq.shift
		WHERE NOT eh.email = sq.email
		ORDER BY carplate ASC, ut.name ASC
  `
	args := pgx.NamedArgs{
		"Email": email,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eventHelper structs.EventHelper
		err := rows.Scan(
			&eventHelper.BusId,
			&eventHelper.Carplate,
			&eventHelper.Name,
			&eventHelper.Email,
			&eventHelper.Shift,
		)
		if err != nil {
			return nil, err
		}
		eventHelpers = append(eventHelpers, eventHelper)
	}
	return eventHelpers, nil
}

func CreateSwapRequest(swap_request structs.SwapRequest) error {
	query := `INSERT INTO swap_request("from", "with", accepted) VALUES (@From, @With, FALSE)`

	args := pgx.NamedArgs{
		"From": swap_request.From,
		"With": swap_request.With,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetSwapRequests(email string) ([]structs.SwapRequestResponse, error) {
	var swap_requests []structs.SwapRequestResponse

	query := `
		WITH sr AS (
			SELECT timestamp, "from", "with", shift FROM swap_request sr JOIN event_helper eh ON sr."with" = eh.email WHERE accepted = FALSE AND ("from" = @Email OR "with" = @Email)
		)
		SELECT timestamp, "from", (SELECT name FROM user_table WHERE email = "from"), "with", (SELECT name FROM user_table WHERE email = "with"), shift FROM sr
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data structs.SwapRequestResponse
		err := rows.Scan(&data.Timestamp, &data.From, &data.FromName, &data.With, &data.WithName, &data.TargetShift)
		if err != nil {
			return nil, err
		}
		swap_requests = append(swap_requests, data)
	}

	return swap_requests, nil
}

func AcceptSwapRequest(swap_request structs.SwapRequest) error {
	updateFrom := `UPDATE event_helper SET shift = NOT shift WHERE email = @From`
	updateWith := `UPDATE event_helper SET shift = NOT shift WHERE email = @With`
	query := `UPDATE swap_request SET accepted = TRUE WHERE "from" = @From AND "with" = @With`

	args := pgx.NamedArgs{
		"From": swap_request.From,
		"With": swap_request.With,
	}

	_, err := config.Dbpool.Exec(context.Background(), updateFrom, args)
	if err != nil {
		return err
	}

	_, err = config.Dbpool.Exec(context.Background(), updateWith, args)
	if err != nil {
		return err
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSwapRequest(swap_request structs.SwapRequest) error {
	query := `DELETE FROM swap_request WHERE "from" = @From AND "with" = @With`

	args := pgx.NamedArgs{
		"From": swap_request.From,
		"With": swap_request.With,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func GetAcceptedSwapRequests() ([]structs.SwapRequestResponse, error) {
	var swap_requests []structs.SwapRequestResponse

	query := `
		WITH sr AS (
			SELECT timestamp, "from", "with", shift FROM swap_request sr JOIN event_helper eh ON sr."with" = eh.email WHERE accepted = TRUE
		)
		SELECT timestamp, "from", (SELECT name FROM user_table WHERE email = "from"), "with", (SELECT name FROM user_table WHERE email = "with"), shift FROM sr
		ORDER BY timestamp DESC 
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data structs.SwapRequestResponse
		err := rows.Scan(&data.Timestamp, &data.From, &data.FromName, &data.With, &data.WithName, &data.TargetShift)
		if err != nil {
			return nil, err
		}
		swap_requests = append(swap_requests, data)
	}

	return swap_requests, nil
}
