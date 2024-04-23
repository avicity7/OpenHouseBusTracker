package services

import (
	"context"
	"server/config"
	"server/structs"

	// "github.com/jackc/pgx/v4"
)

func GetSchedule() ([]structs.Schedule, error) {
	var schedules []structs.Schedule

	query := `
		SELECT 
			b."BusId",
			b."Carplate" AS Bus_Carplate,
			r."Description" AS Route_Description,
			d."Name" AS Driver_Name,
			bs."StartTime",
			bs."EndTime"
		FROM 
			"BusSchedule" bs
		JOIN 
			"Bus" b ON bs."BusId" = b."BusId"
		JOIN 
			"Route" r ON bs."RouteId" = r."RouteId"
		JOIN 
			"Driver" d ON bs."AssignedDriver" = d."DriverId"
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
			&schedule.RouteDescription,
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