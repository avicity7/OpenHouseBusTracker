package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateBus(carplate string) error {
	query := `INSERT INTO bus(carplate, status) VALUES (@Carplate, FALSE)`

	args := pgx.NamedArgs{
		"Carplate": carplate,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}

func GetBuses() ([]structs.EventBus, error) {
	var output []structs.EventBus

	query := `
		SELECT * FROM bus
		ORDER BY carplate DESC
	`
	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.EventBus{}, err
	}

	for rows.Next() {
		var eventBus structs.EventBus
		rows.Scan(&eventBus.Carplate, &eventBus.Status)
		output = append(output, eventBus)
	}

	return output, nil
}

func GetBusStatus(carplate string) (bool, error) {
	var status bool
	query := `SELECT status FROM bus WHERE carplate = @Carplate`
	args := pgx.NamedArgs{
		"Carplate": carplate,
	}
	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&status)
	if err != nil {
		fmt.Println("Error query status", err)
		return false, err
	}

	return status, nil
}

func DeleteBus(carplate string) error {
	query := `DELETE FROM bus WHERE carplate = @Carplate`

	args := pgx.NamedArgs{
		"Carplate": carplate,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}

func UpdateBusStatus(status bool, carplate string) error {
	query := `
		UPDATE bus
		SET status = @Status 
		WHERE carplate = @Carplate
	`

	args := pgx.NamedArgs{
		"Carplate": carplate,
		"Status":   status,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}
