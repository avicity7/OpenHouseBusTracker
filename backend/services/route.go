package services

import (
	"context"
	"database/sql"
	"server/config"
	"server/structs"
	// "github.com/jackc/pgx/v5"
)

// Query to fetch all routes
func GetAllRoutes() ([]structs.Route, error) {
	rows, err := config.Dbpool.Query(context.Background(), "SELECT route_name FROM route")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []structs.Route
	for rows.Next() {
		var route structs.Route
		if err := rows.Scan(&route.RouteName); err != nil {
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
func GetRouteByName(routeName string) (*structs.Route, error) {
	var route structs.Route
	err := config.Dbpool.QueryRow(context.Background(), "SELECT route_name FROM route WHERE route_name = $1", routeName).Scan(&route.RouteName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &route, nil
}
