package techniqueStructs

import "houseGoProject/structs"

type Laptop struct {
	ID          int
	Technique   structs.Techniques
	Name        string
	Diagonal    int
	GraphicCard string
	Processor   string
}
