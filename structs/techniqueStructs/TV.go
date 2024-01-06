package techniqueStructs

import "houseGoProject/structs"

type TV struct {
	ID        int
	technique structs.Techniques
	name      string
	diagonal  int
	isSmart   bool
}
