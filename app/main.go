package main

import (
	"fmt"
	"go-games-api/app/models"
	"go-games-api/db"
	"go-games-api/utilities"
	"log"
)

func main() {
	conn := db.Connect()
	defer conn.Close()
	query := "SELECT id,Total,NumTurns,created_at,updated_at,COALESCE(user_id,0) FROM yachts"
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var yachts []models.Yacht
	for rows.Next() {
		var yacht models.Yacht
		err := rows.Scan(&yacht.Total, &yacht.NumTurns, &yacht.Id, &yacht.CreatedAt, &yacht.UpdatedAt, &yacht.UserId)
		if err != nil {
			log.Fatal(err)
		}
		yachts = append(yachts, yacht)
	}
	json, err := utilities.PrettyStruct(yachts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(json)
}
