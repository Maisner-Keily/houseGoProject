package houseFunctions

import (
	"context"
	"fmt"
	"houseGoProject/structs/furnitureStructs"
	"os"
)

func getRoomBeds(ctx context.Context, roomId int) []furnitureStructs.Bed {
	conn := getConn(ctx)
	var beds []furnitureStructs.Bed
	bedsRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM bed WHERE roomid=", roomId))

	for bedsRows.Next() {
		var bedID int
		var furniture furnitureStructs.Furniture
		var name, mattress string
		var roomID int

		err := bedsRows.Scan(&bedID, &furniture, &name, &mattress, &roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan beds row data: %v\n", err)
			os.Exit(1)
		}

		bed := furnitureStructs.Bed{
			ID:        bedID,
			Name:      name,
			Furniture: furniture,
			Mattress:  mattress,
		}

		beds = append(beds, bed)
	}

	return beds
}

func getRoomBedsideTables(ctx context.Context, roomId int) []furnitureStructs.BedsideTable {
	conn := getConn(ctx)
	var bedsideTables []furnitureStructs.BedsideTable
	bedsideTablesRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM bedsidetable WHERE roomid=", roomId))

	for bedsideTablesRows.Next() {
		var bedsideTableID int
		var furniture furnitureStructs.Furniture
		var name string
		var boxCount, roomID int

		err := bedsideTablesRows.Scan(&bedsideTableID, &furniture, &name, &boxCount, &roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan bedsideTable row data: %v\n", err)
			os.Exit(1)
		}

		bedsideTable := furnitureStructs.BedsideTable{
			ID:        bedsideTableID,
			Name:      name,
			Furniture: furniture,
		}

		bedsideTables = append(bedsideTables, bedsideTable)
	}

	return bedsideTables
}
