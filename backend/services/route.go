package services

import (
	"context"
	"database/sql"
	"fmt"
	"server/config"
	"server/structs"
	// "github.com/jackc/pgx/v5"
)

func CreateRoute(route structs.Route) error {
	_, err := config.Dbpool.Exec(context.Background(), `INSERT INTO route (route_name, color) VALUES ($1, $2)`, route.RouteName, route.Color)
	if err != nil {
		fmt.Println("Error creating route:", err)
		return err
	}
	return nil
}

// Query to fetch all routes
func GetAllRoutes() ([]structs.Route, error) {
	rows, err := config.Dbpool.Query(context.Background(), "SELECT * FROM route")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []structs.Route
	for rows.Next() {
		var route structs.Route
		if err := rows.Scan(&route.RouteName, &route.Color); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}

// retrieves a route by its name from the Dbpool
func GetRouteByName(routeName string) (structs.Route, error) {
	var route structs.Route
	err := config.Dbpool.QueryRow(context.Background(), "SELECT route_name FROM route WHERE route_name = $1", routeName).Scan(&route.RouteName)
	if err == sql.ErrNoRows {
		return structs.Route{}, nil
	} else if err != nil {
		return structs.Route{}, err
	}

	return route, nil
}

// update route
func UpdateRoute(routeName string) error {
	_, err := config.Dbpool.Exec(context.Background(), `UPDATE route SET route_name = $1 WHERE route_name = $2`, routeName)
	if err != nil {
		fmt.Println("Error updating route step:", err)
		return err
	}
	return nil
}

// delete route
func DeleteRoute(routeName string) error {
	_, err := config.Dbpool.Exec(context.Background(), `DELETE FROM route WHERE route_name = $1`, routeName)
	if err != nil {
		fmt.Println("Error deleting route step:", err)
		return err
	}
	return nil
}
