package services

import (
	"context"
	"database/sql"
	"fmt"
	"server/config"
	"server/structs"
	// "github.com/jackc/pgx/v5"
)

func GetAllRouteSteps(routeName string) ([]structs.RouteStep, error) {
	rows, err := config.Dbpool.Query(context.Background(), `SELECT route_name, stop_name, "order" FROM route_step WHERE route_name = $1 ORDER BY "order" ASC`, routeName)
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

// get route
func GetRouteStep(routeName, stopName string) (*structs.RouteStep, error) {
	var routeStep structs.RouteStep
	err := config.Dbpool.QueryRow(context.Background(), `SELECT route_name, stop_name, "order" FROM route_step WHERE route_name = $1 AND stop_name = $2`, routeName, stopName).Scan(&routeStep.RouteName, &routeStep.StopName, &routeStep.Order)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		fmt.Println("Error fetching route step:", err)
		return nil, err
	}

	return &routeStep, nil
}

// create route
func CreateRouteStep(routeStep *structs.RouteStep) error {
	_, err := config.Dbpool.Exec(context.Background(), `INSERT INTO route_step (route_name, stop_name, "order") VALUES ($1, $2, $3)`, routeStep.RouteName, routeStep.StopName, routeStep.Order)
	if err != nil {
		fmt.Println("Error creating route step:", err)
		return err
	}
	return nil
}

// update route
func UpdateRouteStep(routeStep *structs.RouteStep) error {
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
