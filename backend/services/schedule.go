package services

import (
	"context"
	"fmt"
	"server/config"
	"server/structs"
	// "github.com/jackc/pgx/v5"
)

func GetSchedule() ([]structs.Schedule, error) {
	var schedules []structs.Schedule

	query := `
		SELECT 
			bs.bus_schedule_id,
			b.carplate AS Bus_Carplate,
			r.route_name AS Route_Name,
			d.driver_name AS Driver_Name,
			bs.start_time,
			bs.end_time
		FROM 
			bus_schedule bs
		JOIN 
			bus b ON bs.carplate = b.carplate
		JOIN 
			route r ON bs.route_name = r.route_name
		JOIN 
			driver d ON bs.driver_id = d.driver_id
    `

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule structs.Schedule
		err := rows.Scan(
			&schedule.BusScheduleId,
			&schedule.Carplate,
			&schedule.RouteName,
			&schedule.DriverName,
			&schedule.StartTime,
			&schedule.EndTime,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func CreateBusSchedule(schedule structs.NewSchedule) error {
    query := `
		INSERT INTO bus_schedule (carplate, route_name, driver_id, start_time, end_time) 
		VALUES ($1, $2, $3, $4, $5)
    `
	_, err := config.Dbpool.Exec(context.Background(), query, 
		schedule.Carplate, 
		schedule.RouteName, 
		schedule.DriverId, 
		schedule.StartTime, 
		schedule.EndTime,
	)

	if err != nil {
		fmt.Println("Error inserting schedule:", err)
		return err
	}
	
    return nil
}

func UpdateBusSchedule(schedule structs.UpdateSchedule) error {
    query := `
		UPDATE bus_schedule
		SET 
			carplate = $1,
			route_name = $2,
			driver_id = $3,
			start_time = $4,
			end_time = $5
		WHERE
			bus_schedule_id = $6
    `
	_, err := config.Dbpool.Exec(context.Background(), query, 
		schedule.Carplate, 
		schedule.RouteName, 
		schedule.DriverId, 
		schedule.StartTime, 
		schedule.EndTime,
		schedule.BusScheduleId,
	)

	if err != nil {
		fmt.Println("Error updating schedule:", err)
		return err
	}
	
    return nil
}

func DeleteBusSchedule(scheduleID int) error {
    query := `
		DELETE FROM bus_schedule
		WHERE bus_schedule_id = $1
    `
	_, err := config.Dbpool.Exec(context.Background(), query, scheduleID)

	if err != nil {
		fmt.Println("Error deleting schedule:", err)
		return err
	}
	
    return nil
}

func GetDropdownData() ([]structs.ScheduleDropdownData, error) {
	var dropdownData []structs.ScheduleDropdownData

	query := `
		SELECT 
			b.carplate,
			r.route_name,
			d.driver_name,
			d.driver_id
		FROM 
			bus b,
			route r,
			driver d
    `

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data structs.ScheduleDropdownData
		var driver structs.Driver
		err := rows.Scan(
			&data.Carplate,
			&data.RouteName,
			&driver.DriverName,
			&driver.DriverId,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		data.Driver = append(data.Driver, driver)
		dropdownData = append(dropdownData, data)

	}

	return dropdownData, nil
}

func GetScheduleByID(id int) (structs.Schedule, error) {
	var schedule structs.Schedule

	fmt.Println("Received schedule ID in service:", id) // Add this line to log the received ID

	query := `
		SELECT 
			bs.bus_schedule_id,
			b.carplate AS Bus_Carplate,
			r.route_name AS Route_Name,
			d.driver_name AS Driver_Name,
			bs.start_time,
			bs.end_time
		FROM 
			bus_schedule bs
		JOIN 
			bus b ON bs.carplate = b.carplate
		JOIN 
			route r ON bs.route_name = r.route_name
		JOIN 
			driver d ON bs.driver_id = d.driver_id
		WHERE
			bs.bus_schedule_id = $1
    `

	err := config.Dbpool.QueryRow(context.Background(), query, id).Scan(
		&schedule.BusScheduleId,
		&schedule.Carplate,
		&schedule.RouteName,
		&schedule.DriverName,
		&schedule.StartTime,
		&schedule.EndTime,
	)
	if err != nil {
		return structs.Schedule{}, err
	}

	return schedule, nil
}
