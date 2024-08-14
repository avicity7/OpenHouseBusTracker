package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateFollowBus(bus_id string, email string) error {
	query := `
		INSERT INTO event_helper (bus_id, email)
		VALUES (@BusId, @Email)
	`

	args := pgx.NamedArgs{
		"BusId": bus_id,
		"Email": email,
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
		SELECT bus_schedule_id, eh.bus_id, b.carplate, driver_name, route_name, bs.start_time, bs.end_time FROM event_helper eh
		JOIN bus b ON eh.bus_id = b.bus_id 
		JOIN bus_schedule bs ON eh.bus_id = bs.bus_id 
		JOIN driver d ON bs.driver_id = d.driver_id 
		WHERE email = @Email
		AND shift = (TO_CHAR(NOW() AT TIME ZONE 'Etc/GMT-8', 'HH24:MI:SS') >= '00:00:00') AND (TO_CHAR(NOW() AT TIME ZONE 'Etc/GMT-8', 'HH24:MI:SS') <= '14:00:00')
		AND NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
		ORDER BY shift DESC
		LIMIT 1
	`

	args := pgx.NamedArgs{
		"Email": email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&eventSchedule.ScheduleId, &eventSchedule.BusId, &eventSchedule.Carplate, &eventSchedule.DriverName, &eventSchedule.RouteName, &eventSchedule.BusStartTime, &eventSchedule.BusEndTime)
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
		SELECT bs.bus_schedule_id, bs.bus_id, b.carplate, d.driver_name, route_name, email, bs.start_time AS "bus_start_time", bs.end_time AS "bus_end_time" FROM bus_schedule bs 
		JOIN bus b ON bs.bus_id = b.bus_id
		JOIN event_helper eh ON bs.bus_id = eh.bus_id
		JOIN driver d ON bs.driver_id = d.driver_id
		WHERE NOW() AT TIME ZONE 'Etc/GMT-8' BETWEEN bs.start_time AND bs.end_time
	`

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.FollowBusEvent{}, err
	}

	for rows.Next() {
		var followBus structs.FollowBusEvent
		rows.Scan(&followBus.ScheduleId, &followBus.BusId, &followBus.Carplate, &followBus.DriverName, &followBus.RouteName, &followBus.Email, &followBus.BusStartTime, &followBus.BusEndTime)
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

func GetEvents(scheduleId int) ([]structs.Event, error) {
	var output []structs.Event

	query := `
		SELECT rs.stop_name, rs."order", event_id, "timestamp" FROM bus_schedule bs
		JOIN "event" e ON e.bus_id = bs.bus_id 
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

func CreateEvent(bus_id string, routeName string, eventId int, stopName string) error {
	query := `
		INSERT INTO event("timestamp", bus_id, route_name, event_id, stop_name)
		VALUES (NOW() AT TIME ZONE 'Etc/GMT-8', @BusId, @RouteName, @EventId, @StopName) 
	`

	args := pgx.NamedArgs{
		"BusId":     bus_id,
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

func DeleteLastEvent(bus_id string) error {
	query := `
		DELETE FROM "event" 
		WHERE timestamp IN (SELECT timestamp FROM "event" WHERE bus_id = @BusId ORDER BY timestamp DESC LIMIT 1)
		AND bus_id IN (SELECT bus_id FROM "event" WHERE bus_id = @BusId ORDER BY timestamp DESC LIMIT 1)
	`

	args := pgx.NamedArgs{
		"BusId": bus_id,
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
			SELECT *, ROW_NUMBER() OVER (PARTITION BY bus_id ORDER BY timestamp DESC) row_num FROM "event" e
		)
		SELECT bus.bus_id, bus.carplate, r.route_name, color, et.event_name, lng, lat, "timestamp" FROM ranked r
		JOIN bus ON bus.bus_id = r.bus_id
		JOIN stop ON stop.stop_name = r.stop_name
		JOIN route ON route.route_name = r.route_name
		JOIN event_type et ON et.event_id = r.event_id
		WHERE row_num = 1 
		AND bus.bus_id IN (
			SELECT bus_id FROM bus_schedule bs 
			WHERE end_time >= NOW() AT TIME ZONE 'Etc/GMT-8'
			AND start_time <= NOW() AT TIME ZONE 'Etc/GMT-8'
		)
	`

	var currentBuses []structs.CurrentBus

	rows, err := config.Dbpool.Query(context.Background(), query)
	if err != nil {
		return []structs.CurrentBus{}, err
	}

	for rows.Next() {
		var currentBus structs.CurrentBus
		rows.Scan(&currentBus.BusId, &currentBus.Carplate, &currentBus.RouteName, &currentBus.Color, &currentBus.EventType, &currentBus.Lng, &currentBus.Lat, &currentBus.Timestamp)
		currentBuses = append(currentBuses, currentBus)
	}

	return currentBuses, nil
}
