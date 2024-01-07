package techniqueStructs

import "houseGoProject/structs"

type TV struct {
	ID        int
	Technique structs.Techniques
	Name      string
	Diagonal  int
	IsSmart   bool
}
