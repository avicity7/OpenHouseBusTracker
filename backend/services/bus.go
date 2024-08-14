package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateBus(carplate string) error {
	query := `INSERT INTO bus(bus_id, carplate, status) VALUES (gen_random_uuid(), @Carplate, FALSE)`

	args := pgx.NamedArgs{
		"Carplate": carplate,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	return nil
}

func RefreshBuses() error {
	query := `
	UPDATE bus 
	SET status = false 
	WHERE bus_id NOT IN 
		(
			SELECT bus.bus_id FROM bus
			JOIN bus_schedule bs ON bs.bus_id = bus.bus_id 
			WHERE bs.end_time > NOW() AT TIME ZONE 'Etc/GMT-8'
		)
	`

	_, err := config.Dbpool.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func GetBuses() ([]structs.EventBus, error) {
	var output []structs.EventBus

	query := `
		SELECT * FROM bus
		ORDER BY carplate ASC
	`
	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.EventBus{}, err
	}

	for rows.Next() {
		var eventBus structs.EventBus
		rows.Scan(&eventBus.Carplate, &eventBus.Status, &eventBus.Hidden, &eventBus.BusId)
		output = append(output, eventBus)
	}

	return output, nil
}

func GetBusStatus(bus_id string) (bool, error) {
	var status bool
	query := `SELECT status FROM bus WHERE bus_id = @BusId`
	args := pgx.NamedArgs{
		"BusId": bus_id,
	}
	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&status)
	if err != nil {
		fmt.Println("Error query status", err)
		return false, err
	}

	return status, nil
}

func DeleteBus(bus_id string) error {
	query := `DELETE FROM bus WHERE bus_id = @BusId`

	args := pgx.NamedArgs{
		"BusId": bus_id,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}

func UpdateBus(bus_id string, newCarplate string) error {
	query := `
			UPDATE bus
      SET carplate = $1
			WHERE bus_id = $2
		`
	_, err := config.Dbpool.Exec(context.Background(), query, newCarplate, bus_id)
	if err != nil {
		fmt.Println("Error updating bus carplate:", err)
		return err
	}
	fmt.Println("Bus carplate updated successfully")
	return nil
}

func UpdateBusStatus(status bool, bus_id string) error {
	query := `
		UPDATE bus
		SET status = @Status 
		WHERE bus_id = @BusId
	`

	args := pgx.NamedArgs{
		"BusId":  bus_id,
		"Status": status,
	}
	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}

func UpdateBusVisibility(hidden string, bus_id string) error {
	query := `
		UPDATE bus
		SET hidden = @Hidden 
		WHERE bus_id = @BusId
	`

	args := pgx.NamedArgs{
		"BusId":  bus_id,
		"Hidden": hidden,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("error adding driver:", err)
		return err
	}
	return nil
}
