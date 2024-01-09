package houseFunctions

import (
	"context"
	"fmt"
	"houseGoProject/enums/RoomTypes"
	"houseGoProject/structs"
	"os"
)

func getFullHouseRooms(ctx context.Context, houseId int) []structs.Room {
	houseRooms := getHouseRooms(ctx, houseId)
	for index, room := range houseRooms {
		room.Beds = getRoomBeds(ctx, room.ID)
		room.BedsideTables = getRoomBedsideTables(ctx, room.ID)
		room.TVs = getRoomTVs(ctx, room.ID)
		room.Laptops = getRoomLaptops(ctx, room.ID)
		room.WashingMachines = getRoomWashingMachines(ctx, room.ID)

		houseRooms[index] = room
	}

	return houseRooms
}

func getHouseRooms(ctx context.Context, houseId int) []structs.Room {
	conn := getConn(ctx)
	var rooms []structs.Room
	roomRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM room WHERE houseid=", houseId))

	for roomRows.Next() {
		var roomID int
		var roomSizes structs.Sizes
		var roomType RoomTypes.RoomType
		var houseId int

		err := roomRows.Scan(&roomID, &roomSizes, &roomType, &houseId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan room data: %v\n", err)
			os.Exit(1)
		}

		room := structs.Room{
			ID:              roomID,
			Sizes:           roomSizes,
			RoomType:        roomType,
			Beds:            nil,
			BedsideTables:   nil,
			TVs:             nil,
			Laptops:         nil,
			WashingMachines: nil,
		}

		rooms = append(rooms, room)
	}

	return rooms
}
