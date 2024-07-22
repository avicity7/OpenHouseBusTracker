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
			bus b ON bs.bus_id = b.bus_id
		JOIN 
			route r ON bs.route_name = r.route_name
		JOIN 
			driver d ON bs.driver_id = d.driver_id
		ORDER BY
			b.carplate ASC
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
	var existingCount int
	checkQuery := `
        SELECT COUNT(*) FROM bus_schedule 
        WHERE (bus_id = $1 OR driver_id = $2)
    `
	err := config.Dbpool.QueryRow(context.Background(), checkQuery,
		schedule.BusId,
		schedule.DriverId).Scan(&existingCount)

	if err != nil {
		fmt.Println("Error checking existing schedule:", err)
		return err
	}

	if existingCount > 0 {
		return fmt.Errorf("a schedule with the same carplate or driver already exists")
	}

	insertQuery := `
        INSERT INTO bus_schedule (bus_id, route_name, driver_id, start_time, end_time) 
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err = config.Dbpool.Exec(context.Background(), insertQuery,
		schedule.BusId,
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
			bus_id = $1,
			route_name = $2,
			driver_id = $3,
			start_time = $4,
			end_time = $5
		WHERE
			bus_schedule_id = $6
    `
	_, err := config.Dbpool.Exec(context.Background(), query,
		schedule.BusId,
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

func UpdateScheduleRoutes(assignments []structs.UpdateScheduleRoute) error {
	for _, assignment := range assignments {
		query := `
			UPDATE bus_schedule
			SET route_name = $1
			WHERE bus_id = $2
		`
		_, err := config.Dbpool.Exec(context.Background(), query, assignment.RouteName, assignment.BusId)
		if err != nil {
			fmt.Println("Error updating schedule:", err)
			return err
		}
	}
	return nil
}

func DeleteBusSchedule(scheduleID []int) error {
	query := `
		DELETE FROM bus_schedule
		WHERE bus_schedule_id = ANY ($1)
    `
	_, err := config.Dbpool.Exec(context.Background(), query, scheduleID)
	if err != nil {
		fmt.Println("Error deleting schedule:", err)
		return err
	}

	return nil
}

// func GetDropdownData() ([]structs.ScheduleDropdownData, error) {
// 	var dropdownData []structs.ScheduleDropdownData

// 	query := `
// 			WITH available_buses AS (
// 			SELECT carplate, bus_id
// 			FROM bus
// 			WHERE bus_id NOT IN (SELECT bus_id FROM bus_schedule)
// 		),
// 		available_drivers AS (
// 			SELECT driver_id, driver_name
// 			FROM driver
// 			WHERE driver_id NOT IN (SELECT driver_id FROM bus_schedule)
// 		)
// 		SELECT 
// 			b.bus_id,
// 			COALESCE(b.carplate, NULL) AS carplate,
// 			r.route_name,
// 			COALESCE(d.driver_name, NULL) AS driver_name,
// 			COALESCE(d.driver_id, NULL) AS driver_id
// 		FROM 
// 			(
// 				SELECT 1 AS dummy
// 			) dummy_table
// 		LEFT JOIN 
// 			available_buses ab ON true
// 		LEFT JOIN 
// 			bus b ON ab.bus_id = b.bus_id
// 		LEFT JOIN 
// 			route r ON 1=1
// 		LEFT JOIN 
// 			available_drivers d ON true
// 		ORDER BY	
// 			ab.carplate ASC, d.driver_id ASC;
//     `

// 	rows, err := config.Dbpool.Query(context.Background(), query)
// 	if err != nil {
// 		fmt.Println("Error executing query:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var data structs.ScheduleDropdownData
// 		var driver structs.Driver
// 		err := rows.Scan(
// 			&data.BusId,
// 			&data.Carplate,
// 			&data.RouteName,
// 			&driver.DriverName,
// 			&driver.DriverId,
// 		)
// 		if err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			return nil, err
// 		}

// 		data.Driver = append(data.Driver, driver)
// 		dropdownData = append(dropdownData, data)

// 	}

// 	return dropdownData, nil
// }

func GetDropdownData() (structs.ScheduleDropdownData, error) {
	var dropdownData structs.ScheduleDropdownData

	busQuery := `SELECT bus_id, carplate FROM bus 
				 WHERE bus_id 
				 NOT IN (SELECT bus_id FROM bus_schedule)`
	busRows, err := config.Dbpool.Query(context.Background(), busQuery)
	if err != nil {
		fmt.Println("Error executing bus query:", err)
		return dropdownData, err
	}
	defer busRows.Close()

	var buses []structs.EventBus
	for busRows.Next() {
		var bus structs.EventBus
		if err := busRows.Scan(&bus.BusId, &bus.Carplate); err != nil {
			fmt.Println("Error scanning bus row:", err)
			return dropdownData, err
		}
		buses = append(buses, bus)
	}
	dropdownData.Buses = buses

	routeQuery := `SELECT route_name, color FROM route`
	routeRows, err := config.Dbpool.Query(context.Background(), routeQuery)
	if err != nil {
		fmt.Println("Error executing route query:", err)
		return dropdownData, err
	}
	defer routeRows.Close()

	var routes []structs.Route
	for routeRows.Next() {
		var route structs.Route
		if err := routeRows.Scan(&route.RouteName, &route.Color); err != nil {
			fmt.Println("Error scanning route row:", err)
			return dropdownData, err
		}
		routes = append(routes, route)
	}
	dropdownData.Routes = routes

	driverQuery := `SELECT driver_id, driver_name FROM driver 
					WHERE driver_id 
					NOT IN (SELECT driver_id FROM bus_schedule)`
	driverRows, err := config.Dbpool.Query(context.Background(), driverQuery)
	if err != nil {
		fmt.Println("Error executing driver query:", err)
		return dropdownData, err
	}
	defer driverRows.Close()

	var drivers []structs.Driver
	for driverRows.Next() {
		var driver structs.Driver
		if err := driverRows.Scan(&driver.DriverId, &driver.DriverName); err != nil {
			fmt.Println("Error scanning driver row:", err)
			return dropdownData, err
		}
		drivers = append(drivers, driver)
	}
	dropdownData.Drivers = drivers

	return dropdownData, nil
}

func GetScheduleByID(id int) (structs.UpdateSchedule, error) {
	var schedule structs.UpdateSchedule

	query := `
	SELECT 
		bs.bus_schedule_id,
		b.bus_id,
		b.carplate AS Bus_Carplate,
		r.route_name AS Route_Name,
		d.driver_id AS Driver_Id,
		bs.start_time,
		bs.end_time
	FROM 
		bus_schedule bs
	JOIN 
		bus b ON bs.bus_id = b.bus_id
	JOIN 
		route r ON bs.route_name = r.route_name
	JOIN 
		driver d ON bs.driver_id = d.driver_id
	WHERE
		bs.bus_schedule_id = $1
    `

	err := config.Dbpool.QueryRow(context.Background(), query, id).Scan(
		&schedule.BusScheduleId,
		&schedule.BusId,
		&schedule.Carplate,
		&schedule.RouteName,
		&schedule.DriverId,
		&schedule.StartTime,
		&schedule.EndTime,
	)
	if err != nil {
		return structs.UpdateSchedule{}, err
	}

	return schedule, nil
}

func GetScheduleByUser(email string) ([]structs.Schedule, error) {
	var schedules []structs.Schedule

	query := `
		SELECT bus_schedule_id, eh.carplate, driver_name, route_name, bs.start_time, bs.end_time FROM event_helper eh
		JOIN bus_schedule bs ON eh.carplate = bs.carplate 
		JOIN driver d ON bs.driver_id = d.driver_id 
		WHERE email = $1
		AND (shift = (NOT (CURRENT_TIME AT TIME ZONE 'Etc/GMT-8' >= '12:00:00')) OR (CURRENT_TIME AT TIME ZONE 'Etc/GMT-8' <= '14:00:00'))
		AND NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
		ORDER BY shift DESC
		LIMIT 1
    `

	rows, err := config.Dbpool.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule structs.Schedule
		err := rows.Scan(
			&schedule.BusScheduleId,
			&schedule.Carplate,
			&schedule.DriverName,
			&schedule.RouteName,
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

func GetFutureScheduleByUser(email string) ([]structs.Schedule, error) {
	var schedules []structs.Schedule

	query := `
		WITH current_shifts AS (
			SELECT 
				bs.bus_schedule_id
			FROM 
				event_helper eh
			JOIN 
				bus_schedule bs ON eh.carplate = bs.carplate 
			JOIN 
				driver d ON bs.driver_id = d.driver_id 
			WHERE 
				email = $1
			AND 
				shift = NOT (CURRENT_TIME AT TIME ZONE 'Etc/GMT-8' >= '12:00:00')
			AND 
				NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
		)
		
		SELECT 
			bus_schedule_id,
			eh.carplate, 
			driver_name, 
			route_name, 
			bs.start_time, 
			bs.end_time 
		FROM 
			event_helper eh
		JOIN 
			bus_schedule bs ON eh.carplate = bs.carplate 
		JOIN 
			driver d ON bs.driver_id = d.driver_id 
		WHERE 
			email = $1
		AND 
			bs.start_time > NOW() AT TIME ZONE 'Etc/GMT-8'
		AND 
			bs.bus_schedule_id NOT IN (SELECT bus_schedule_id FROM current_shifts);
    `

	rows, err := config.Dbpool.Query(context.Background(), query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule structs.Schedule
		err := rows.Scan(
			&schedule.BusScheduleId,
			&schedule.Carplate,
			&schedule.DriverName,
			&schedule.RouteName,
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
