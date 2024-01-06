package furnitureStructs

import (
	"go/types"
	"houseGoProject/structs"
)

type Table struct {
	Furniture structs.Furniture
	Sizes     types.Sizes
	Name      string
}
