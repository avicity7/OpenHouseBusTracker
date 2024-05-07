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
		SELECT bus_schedule_id, eh.carplate, driver_name, route_name, start_time, end_time FROM event_helper eh
		JOIN bus_schedule bs ON eh.carplate = bs.carplate 
		JOIN driver d ON bs.driver_id = d.driver_id 
		WHERE email = 'karlorjalo@gmail.com'
		AND NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN start_time AND end_time
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&eventSchedule.ScheduleId, &eventSchedule.Carplate, &eventSchedule.DriverName, &eventSchedule.RouteName, &eventSchedule.StartTime, &eventSchedule.EndTime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return structs.EventSchedule{}, nil
		}
		return structs.EventSchedule{}, err
	}

	return eventSchedule, nil
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

func GetBuses() ([]structs.EventBus, error) {
	var output []structs.EventBus

	query := `
		SELECT * FROM bus
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

func GetRouteSteps(routeName string) ([]structs.RouteStep, error) {
	var output []structs.RouteStep

	query := `
		SELECT stop_name, "order" FROM route_step rs 
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
		rows.Scan(&routeStep.StopName, &routeStep.Order)
		output = append(output, routeStep)
	}

	return output, nil
}

func GetEvents(scheduleId int) ([]structs.Event, error) {
	var output []structs.Event

	query := `
		SELECT stop_name, event_id, "timestamp"  FROM bus_schedule bs
		JOIN "event" e ON e.carplate = bs.carplate 
		WHERE bs.bus_schedule_id = @ScheduleId
		AND "timestamp" BETWEEN bs.start_time AND bs.end_time 
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
		rows.Scan(&event.StopName, &event.EventId, &event.Timestamp)
		output = append(output, event)
	}

	return output, nil
}
