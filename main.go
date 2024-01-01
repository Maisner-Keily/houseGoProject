package main

import (
	"fmt"
	"houseGoProject/enums/Materials"
	structs "houseGoProject/struct"
)

func main() {
	var x = structs.Furniture{
		Sizes: structs.Sizes{
			Width:  11.0,
			Length: 16.0,
			Height: 19.0,
		},
		Material: Materials.Wood,
	}

	fmt.Println(x.Sizes.Width)
	fmt.Println(x.Sizes.Height)
	fmt.Println(x.Sizes.Length)
	fmt.Println(x.Material)
	//fmt.Println("2131221")
}
