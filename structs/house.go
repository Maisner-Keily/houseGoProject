package structs

type House struct {
	ID          int
	Sizes       Sizes
	country     string
	city        string
	street      string
	numberHouse string
	rooms       []Room
	members     []HouseMember
}
