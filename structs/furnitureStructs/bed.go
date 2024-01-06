package furnitureStructs

import "houseGoProject/structs"

type Bed struct {
	ID        int
	Furniture structs.Furniture
	Name      string
	Mattress  string
}
