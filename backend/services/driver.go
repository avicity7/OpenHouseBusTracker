package services

import (
    "context"
    "fmt"
    "server/config"
    "server/structs"
)

// Add Driver
func AddDriver(driver structs.Driver) error {
    query := `INSERT INTO driver (driver_name) VALUES ($1);`
    _, err := config.Dbpool.Exec(context.Background(), query, driver.Name)
    if err != nil {
        fmt.Println("Error adding driver:", err)
        return err
    }
    return nil
}

// Get Driver
func GetDriver() ([]structs.Driver, error) {
    var drivers []structs.Driver
    rows, err := config.Dbpool.Query(context.Background(), "SELECT * FROM driver")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var driver structs.Driver
        if err := rows.Scan(&driver.DriverID, &driver.Name); err != nil {
            return nil, err
        }
        drivers = append(drivers, driver)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return drivers, nil
}

// Update Driver
func UpdateDriver(driver structs.Driver) error {
    query := `UPDATE driver SET driver_name = $1 WHERE driver_id = $2;`
    _, err := config.Dbpool.Exec(context.Background(), query, driver.Name, driver.DriverID)
    if err != nil {
        fmt.Println("Error updating driver:", err)
        return err
    }
    fmt.Println("Driver updated successfully")
    return nil
}

// Delete Driver
func DeleteDriver(driverID int) error {
    query := `DELETE FROM driver WHERE driver_id = $1;`
    _, err := config.Dbpool.Exec(context.Background(), query, driverID)
    if err != nil {
        fmt.Println("Error deleting driver:", err)
        return err
    }
    return nil
}
