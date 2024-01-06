package techniqueStructs

import "houseGoProject/structs"

type WashingMachine struct {
	ID               int
	technique        structs.Techniques
	name             string
	speed            int
	maxClothesWeight int
}
