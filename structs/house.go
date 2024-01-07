package structs

type House struct {
	ID          int
	Sizes       Sizes
	Country     string
	City        string
	Street      string
	NumberHouse string
	Rooms       []Room
	Members     []HouseMember
}
