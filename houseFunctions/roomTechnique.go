package houseFunctions

import (
	"context"
	"fmt"
	"houseGoProject/structs/techniqueStructs"
	"os"
)

func getRoomTVs(ctx context.Context, roomId int) []techniqueStructs.TV {
	conn := getConn(ctx)
	var TVs []techniqueStructs.TV
	TVsRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM tv WHERE roomid=", roomId))

	for TVsRows.Next() {
		var TVID int
		var technique techniqueStructs.Techniques
		var name string
		var diagonal float32
		var isSmart bool
		var roomID int

		err := TVsRows.Scan(&TVID, &technique, &name, &diagonal, &isSmart, &roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan TVs row data: %v\n", err)
			os.Exit(1)
		}

		TV := techniqueStructs.TV{
			ID:        TVID,
			Name:      name,
			Diagonal:  diagonal,
			IsSmart:   isSmart,
			Technique: technique,
		}

		TVs = append(TVs, TV)
	}

	return TVs
}

func getRoomLaptops(ctx context.Context, roomId int) []techniqueStructs.Laptop {
	conn := getConn(ctx)
	var laptops []techniqueStructs.Laptop
	laptopsRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM laptop WHERE roomid=", roomId))

	for laptopsRows.Next() {
		var laptopID int
		var technique techniqueStructs.Techniques
		var name, graphicCard, processor string
		var diagonal float32
		var roomID int

		err := laptopsRows.Scan(&laptopID, &technique, &name, &diagonal, &graphicCard, &processor, &roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan laptop row data: %v\n", err)
			os.Exit(1)
		}

		laptop := techniqueStructs.Laptop{
			ID:          laptopID,
			Technique:   technique,
			Name:        name,
			Diagonal:    diagonal,
			GraphicCard: graphicCard,
			Processor:   processor,
		}

		laptops = append(laptops, laptop)
	}

	return laptops
}

func getRoomWashingMachines(ctx context.Context, roomId int) []techniqueStructs.WashingMachine {
	conn := getConn(ctx)
	var washingMachines []techniqueStructs.WashingMachine
	washingMachinesRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM washingmachine WHERE roomid=", roomId))

	for washingMachinesRows.Next() {
		var washingMachineID int
		var technique techniqueStructs.Techniques
		var name string
		var speed, roomID, maxClothesWeight int

		err := washingMachinesRows.Scan(&washingMachineID, &technique, &name, &speed, &maxClothesWeight, &roomID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan washingMachine row data: %v\n", err)
			os.Exit(1)
		}

		washingMachine := techniqueStructs.WashingMachine{
			ID:               washingMachineID,
			Technique:        technique,
			Name:             name,
			Speed:            speed,
			MaxClothesWeight: maxClothesWeight,
		}

		washingMachines = append(washingMachines, washingMachine)
	}

	return washingMachines
}
