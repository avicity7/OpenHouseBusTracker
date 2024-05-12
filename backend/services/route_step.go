package services

import (
	"context"
	"database/sql"
	"fmt"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5"
)

func GetAllRouteSteps(routeName string) ([]structs.RouteStep, error) {
	rows, err := config.Dbpool.Query(context.Background(), `SELECT route_name, stop_name, "order" FROM route_step ORDER BY "order" ASC`)
	if err != nil {
		fmt.Println("Error fetching route steps:", err)
		return nil, err
	}
	defer rows.Close()

	var routeSteps []structs.RouteStep
	for rows.Next() {
		var routeStep structs.RouteStep
		if err := rows.Scan(&routeStep.RouteName, &routeStep.StopName, &routeStep.Order); err != nil {
			fmt.Println("Error scanning route step row:", err)
			return nil, err
		}
		routeSteps = append(routeSteps, routeStep)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over route step rows:", err)
		return nil, err
	}

	return routeSteps, nil
}

func GetAllStops() ([]string, error) {
	rows, err := config.Dbpool.Query(context.Background(), `SELECT * FROM stop`)
	if err != nil {
		fmt.Println("Error fetching route steps:", err)
		return nil, err
	}
	defer rows.Close()

	var stops []string
	for rows.Next() {
		var stop string
		if err := rows.Scan(&stop); err != nil {
			fmt.Println("Error scanning route step row:", err)
			return nil, err
		}
		stops = append(stops, stop)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over route step rows:", err)
		return nil, err
	}

	return stops, nil
}

// get route
func GetRouteStep(routeName, stopName string) (structs.RouteStep, error) {
	var routeStep structs.RouteStep
	err := config.Dbpool.QueryRow(context.Background(), `SELECT route_name, stop_name, "order" FROM route_step WHERE route_name = $1 AND stop_name = $2`, routeName, stopName).Scan(&routeStep.RouteName, &routeStep.StopName, &routeStep.Order)
	if err == sql.ErrNoRows {
		return structs.RouteStep{}, nil
	} else if err != nil {
		fmt.Println("Error fetching route step:", err)
		return structs.RouteStep{}, err
	}

	return routeStep, nil
}

// create route
func CreateRouteStep(routeStep structs.RouteStep) error {
	var latest int
	err := config.Dbpool.QueryRow(context.Background(), `SELECT "order" FROM route_step rs WHERE route_name = $1 ORDER BY "order" DESC LIMIT 1`, routeStep.RouteName).Scan(&latest)
	if err != nil {
		if err == pgx.ErrNoRows {
			latest = 0
		} else {
			return err
		}
	}
	_, err = config.Dbpool.Exec(context.Background(), `INSERT INTO route_step (route_name, stop_name, "order") VALUES ($1, $2, $3)`, routeStep.RouteName, routeStep.StopName, latest+1)
	if err != nil {
		fmt.Println("Error creating route step:", err)
		return err
	}
	return nil
}

// update route
func UpdateRouteStep(routeStep structs.RouteStep) error {
	_, err := config.Dbpool.Exec(context.Background(), `UPDATE route_step SET "order" = $1 WHERE route_name = $2 AND stop_name = $3`, routeStep.Order, routeStep.RouteName, routeStep.StopName)
	if err != nil {
		fmt.Println("Error updating route step:", err)
		return err
	}
	return nil
}

// delete route
func DeleteRouteStep(routeName, stopName string) error {
	_, err := config.Dbpool.Exec(context.Background(), `DELETE FROM route_step WHERE route_name = $1 AND stop_name = $2`, routeName, stopName)
	if err != nil {
		fmt.Println("Error deleting route step:", err)
		return err
	}
	return nil

}
