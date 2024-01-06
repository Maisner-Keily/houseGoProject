package structs

import (
	"houseGoProject/structs/furnitureStructs"
	"houseGoProject/structs/techniqueStructs"
)

type Room struct {
	Width           float32
	Length          float32
	Height          float32
	Beds            []furnitureStructs.Bed
	BedsideTables   []furnitureStructs.BedsideTable
	Laptops         []techniqueStructs.Laptop
	TVs             []techniqueStructs.TV
	WashingMachines []techniqueStructs.WashingMachine
}

func (room Room) getSquare() float32 {
	return room.width * room.length
}

func (room Room) getVolume() float32 {
	return room.width * room.length * room.height
}

func (room Room) getPerimeter() float32 {
	return 2 * (room.width + room.length)
}

func createRooms() []Room {
	var rooms []Room

	rooms[0] = Room{width: 5.4, length: 3.4, height: 2.5}
	rooms[1] = Room{width: 3.2, length: 4.1, height: 3}
	rooms[2] = Room{width: 4.7, length: 3.2, height: 2.8}

	return rooms
}
