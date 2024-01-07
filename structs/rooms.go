package structs

import (
	"houseGoProject/enums/RoomTypes"
	"houseGoProject/structs/furnitureStructs"
	"houseGoProject/structs/techniqueStructs"
)

type Room struct {
	ID              int
	Sizes           Sizes
	RoomType        RoomTypes.RoomType
	Beds            []furnitureStructs.Bed
	BedsideTables   []furnitureStructs.BedsideTable
	Laptops         []techniqueStructs.Laptop
	TVs             []techniqueStructs.TV
	WashingMachines []techniqueStructs.WashingMachine
}

//func (room Room) getSquare() float32 {
//	return room.Width * room.Length
//}
//
//func (room Room) getVolume() float32 {
//	return room.Width * room.Length * room.Height
//}
//
//func (room Room) getPerimeter() float32 {
//	return 2 * (room.Width + room.Length)
//}
