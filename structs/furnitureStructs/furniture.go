package furnitureStructs

import (
	"houseGoProject/enums/Materials"
)

type Furniture struct {
	Sizes    FurnitureSizes
	Material Materials.Material
}
