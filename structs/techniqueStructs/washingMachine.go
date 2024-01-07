package techniqueStructs

import "houseGoProject/structs"

type WashingMachine struct {
	ID               int
	Technique        structs.Techniques
	Name             string
	Speed            int
	MaxClothesWeight int
}
