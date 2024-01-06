package structs

import (
	"houseGoProject/enums/FamilyRoles"
	"houseGoProject/enums/Genders"
)

type HouseMember struct {
	ID      int
	Name    string
	Surname string
	Gender  Genders.Gender
	Age     int
	Role    FamilyRoles.FamilyRole
}
