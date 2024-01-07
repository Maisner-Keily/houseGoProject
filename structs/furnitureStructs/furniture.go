package structs

import (
	"houseGoProject/enums/Materials"
)

type Furniture struct {
	Sizes    Sizes
	Material Materials.Material
}
