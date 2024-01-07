package furnitureStructs

import (
	"go/types"
	"houseGoProject/structs"
)

type Table struct {
	ID        int
	Furniture structs.Furniture
	Sizes     types.Sizes
	Name      string
}
