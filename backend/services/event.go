package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateFollowBus(carplate string, email string) error {
	query := `
		INSERT INTO event_helper (carplate, email)
		VALUES (@Carplate, @Email)
	`

	args := pgx.NamedArgs{
		"Carplate": carplate,
		"Email":    email,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func GetFollowBus(email string) (structs.EventSchedule, error) {
	var eventSchedule structs.EventSchedule

	query := `
	SELECT bus_schedule_id, eh.carplate, driver_name, route_name, bs.start_time, bs.end_time FROM event_helper eh
	JOIN bus_schedule bs ON eh.carplate = bs.carplate 
	JOIN driver d ON bs.driver_id = d.driver_id 
	WHERE email = @Email
	AND shift = NOT (CURRENT_TIME AT TIME ZONE 'Etc/GMT-8' >= '12:00:00')
	AND NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&eventSchedule.ScheduleId, &eventSchedule.Carplate, &eventSchedule.DriverName, &eventSchedule.RouteName, &eventSchedule.BusStartTime, &eventSchedule.BusEndTime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return structs.EventSchedule{}, nil
		}
		return structs.EventSchedule{}, err
	}

	return eventSchedule, nil
}

func GetAllFollowBus() ([]structs.FollowBusEvent, error) {
	var output []structs.FollowBusEvent

	query := `
		SELECT bs.bus_schedule_id, bs.carplate, d.driver_name, route_name, email, bs.start_time AS "bus_start_time", bs.end_time AS "bus_end_time", eh.start_time AS "student_start_time", eh.end_time AS "student_end_time" FROM bus_schedule bs 
		JOIN event_helper eh ON bs.carplate = eh.carplate
		JOIN driver d ON bs.driver_id = d.driver_id
		WHERE NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.FollowBusEvent{}, err
	}

	for rows.Next() {
		var followBus structs.FollowBusEvent
		rows.Scan(&followBus.ScheduleId, &followBus.Carplate, &followBus.DriverName, &followBus.RouteName, &followBus.Email, &followBus.BusStartTime, &followBus.BusEndTime, &followBus.StudentStartTime, &followBus.StudentEndTime)
		output = append(output, followBus)
	}

	return output, nil
}

func DeleteFollowBus(email string) error {
	query := `
		DELETE FROM event_helper
		WHERE email = @Email
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func GetRouteSteps(routeName string) ([]structs.RouteStep, error) {
	var output []structs.RouteStep

	query := `
		SELECT * FROM route_step rs 
		WHERE route_name = @RouteName
	`
	args := pgx.NamedArgs{
		"RouteName": routeName,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return []structs.RouteStep{}, err
	}

	for rows.Next() {
		var routeStep structs.RouteStep
		rows.Scan(&routeStep.RouteName, &routeStep.StopName, &routeStep.Order)
		output = append(output, routeStep)
	}

	return output, nil
}

func GetEvents(scheduleId int) ([]structs.Event, error) {
	var output []structs.Event

	query := `
		SELECT rs.stop_name, rs."order", event_id, "timestamp" FROM bus_schedule bs
		JOIN "event" e ON e.carplate = bs.carplate 
		JOIN route_step rs ON rs.route_name = bs.route_name AND rs.stop_name = e.stop_name
		WHERE bs.bus_schedule_id = @ScheduleId
		AND "timestamp" BETWEEN bs.start_time AND bs.end_time 
		ORDER BY "timestamp" DESC
	`
	args := pgx.NamedArgs{
		"ScheduleId": scheduleId,
	}

	rows, err := config.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return []structs.Event{}, err
	}

	for rows.Next() {
		var event structs.Event
		rows.Scan(&event.StopName, &event.Order, &event.EventId, &event.Timestamp)
		output = append(output, event)
	}

	return output, nil
}

func CreateEvent(carplate string, routeName string, eventId int, stopName string) error {
	query := `
		INSERT INTO event("timestamp", carplate, route_name, event_id, stop_name)
		VALUES (NOW() AT TIME ZONE 'Etc/GMT-8', @Carplate, @RouteName, @EventId, @StopName) 
	`

	args := pgx.NamedArgs{
		"Carplate":  carplate,
		"RouteName": routeName,
		"EventId":   eventId,
		"StopName":  stopName,
	}

	_, err := config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	config.Melody.Broadcast([]byte("refresh"))

	return nil
}

func GetCurrentBuses() ([]structs.CurrentBus, error) {
	query := `
		WITH ranked AS (
			SELECT *, ROW_NUMBER() OVER (PARTITION BY carplate ORDER BY timestamp DESC) row_num FROM "event" e
		)
		SELECT carplate, r.route_name, color, et.event_name, lng, lat, "timestamp" FROM ranked r
		JOIN stop ON stop.stop_name = r.stop_name
		JOIN route ON route.route_name = r.route_name
		JOIN event_type et ON et.event_id = r.event_id
		WHERE row_num = 1 
		AND carplate IN (
			SELECT carplate FROM bus_schedule bs 
			WHERE end_time >= NOW() AT TIME ZONE 'Etc/GMT-8'
		)
	`

	var currentBuses []structs.CurrentBus

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.CurrentBus{}, err
	}

	for rows.Next() {
		var currentBus structs.CurrentBus
		rows.Scan(&currentBus.Carplate, &currentBus.RouteName, &currentBus.Color, &currentBus.EventType, &currentBus.Lng, &currentBus.Lat, &currentBus.Timestamp)
		currentBuses = append(currentBuses, currentBus)
	}

	return currentBuses, nil
}
