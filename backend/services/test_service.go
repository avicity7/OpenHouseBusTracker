package services

import (
	"context"
	"encoding/json"
	"fmt"
	"server/config"

	"github.com/jackc/pgx/v5/pgtype"
)

type Result struct {
	Id        string
	Data      float32
	Timestamp pgtype.Timestamp
}

type Results []Result

// interface{} is used here to dictate that the value in the key:value of the returned map can be of ANY type.
func GetTestService() []byte {
	results := make(Results, 0)
	rows, err := config.Dbpool.Query(context.Background(), "SELECT * FROM test_data")
	if err != nil {
		println(err)
		println("unable to query")
	} else {
		for rows.Next() {
			var id string
			var data float32
			var timestamp pgtype.Timestamp
			rows.Scan(&id, &data, &timestamp)
			result := &Result{Id: id, Data: data, Timestamp: timestamp}
			if err != nil {
				fmt.Println("Formatting error")
			}
			results = append(results, *result)
		}
	}
	formatted, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Formatting error")
	}
	header := []byte("data ")
	config.Melody.Broadcast(append(header, formatted...))
	return formatted
}
