package techniqueStructs

import "houseGoProject/structs"

type Laptop struct {
	ID          int
	technique   structs.Techniques
	name        string
	diagonal    int
	graphicCard string
	processor   string
}
