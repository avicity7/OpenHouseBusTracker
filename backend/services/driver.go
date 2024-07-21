package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"
)

// Add Driver
func AddDriver(driver structs.Driver) error {
	query := `INSERT INTO driver (driver_name) VALUES ($1);`
	_, err := config.Dbpool.Exec(context.Background(), query, driver.DriverName)
	if err != nil {
		fmt.Println("Error adding driver:", err)
		return err
	}
	return nil
}

func GetDrivers() ([]structs.Driver, error) {
	var drivers []structs.Driver
	rows, err := config.Dbpool.Query(context.Background(), "SELECT * FROM driver")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var driver structs.Driver
		if err := rows.Scan(&driver.DriverId, &driver.DriverName); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return drivers, nil
}

// Update Driver
func UpdateDriver(driver structs.Driver) error {
	query := `UPDATE driver SET driver_name = $1 WHERE driver_id = $2;`
	_, err := config.Dbpool.Exec(context.Background(), query, driver.DriverName, driver.DriverId)
	if err != nil {
		fmt.Println("Error updating driver:", err)
		return err
	}
	fmt.Println("Driver updated successfully")
	return nil
}

// Delete Driver
func DeleteDriver(driverID int) error {
	query := `DELETE FROM driver WHERE driver_id = $1;`
	_, err := config.Dbpool.Exec(context.Background(), query, driverID)
	if err != nil {
		fmt.Println("Error deleting driver:", err)
		return err
	}
	return nil
}

func GetScheduleTimeDiff() ([]structs.TimeDiff, error) {
	var driversTime []structs.TimeDiff

	query := `
        SELECT driver_id, start_time, end_time, (end_time - start_time) AS time_diff FROM bus_schedule
    `

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var td structs.TimeDiff
		err := rows.Scan(
			&td.DriverId,
			&td.StartTime,
			&td.EndTime,
			&td.TimeDifference,
		)
		if err != nil {
			return nil, err
		}

		driversTime = append(driversTime, td)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return driversTime, nil
}
