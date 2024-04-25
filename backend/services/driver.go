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
func GetDriver(driverID int) (structs.Driver, error) {
    var driver structs.Driver
    query := `SELECT * FROM driver WHERE driver_id = $1;`
    err := config.Dbpool.QueryRow(context.Background(), query, driverID).Scan(&driver.DriverID, &driver.Name)
    if err != nil {
        fmt.Println("Error getting driver:", err)
        return driver, err
    }
    return driver, nil
}

// Update Driver
func UpdateDriver(driver structs.Driver) error {
    query := `UPDATE driver SET driver_name = $1 WHERE driver_id = $2;`
    _, err := config.Dbpool.Exec(context.Background(), query, driver.Name, driver.DriverID)
    if err != nil {
        fmt.Println("Error updating driver:", err)
        return err
    }
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
