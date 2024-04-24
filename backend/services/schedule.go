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
			&schedule.BusId,
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

// to be confirmed, might need unique identifier for BusSchedule for Update and Delete 
// below is an example of checking composite keys to update

// func UpdateBusSchedule(schedule structs.NewSchedule) error {
//     query := `
//         UPDATE "BusSchedule"
//         SET "AssignedDriver" = $1, "StartTime" = $2, "EndTime" = $3
//         WHERE "BusId" = $4 AND "RouteId" = $5 AND "AssignedDriver" = $6 AND "StartTime" = $7 AND "EndTime" = $8
//     `

//     _, err := config.Dbpool.Exec(context.Background(), query, schedule.AssignedDriver, schedule.StartTime, schedule.EndTime, schedule.BusId, schedule.RouteId, schedule.AssignedDriver, schedule.StartTime, schedule.EndTime)
//     if err != nil {
//         return err
//     }

//     return nil
// }