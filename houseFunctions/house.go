package houseFunctions

import (
	"context"
	"encoding/json"
	"fmt"
	"houseGoProject/structs"
	"log"
	"os"
)

func PrintHouse() {
	house := getFullHouseData()

	empJSON, err := json.MarshalIndent(house, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("House %s\n", string(empJSON))
}

func getFullHouseData() structs.House {
	ctx := context.Background()
	conn := getConn(ctx)
	defer conn.Close(ctx)

	house := getHouse(ctx)
	house.Members = getHouseMembers(ctx, house.ID)
	house.Rooms = getFullHouseRooms(ctx, house.ID)

	return house
}

func getHouse(ctx context.Context) structs.House {
	conn := getConn(ctx)
	var house structs.House
	houseRow := conn.QueryRow(ctx, "SELECT * FROM house limit 1")

	var ID int
	var Sizes structs.Sizes
	var Country, City, Street, NumberHouse string

	err := houseRow.Scan(&ID, &Sizes, &Country, &City, &Street, &NumberHouse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to scan house data: %v\n", err)
		os.Exit(1)
	}

	house = structs.House{
		ID:          ID,
		Sizes:       Sizes,
		Country:     Country,
		City:        City,
		Street:      Street,
		NumberHouse: NumberHouse,
		Rooms:       nil,
		Members:     nil,
	}

	return house
}
