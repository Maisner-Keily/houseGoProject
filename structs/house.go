package structs

type House struct {
	ID          int
	country     string
	city        string
	street      string
	numberHouse string
	rooms       []Room
	members     []HouseMember
}
