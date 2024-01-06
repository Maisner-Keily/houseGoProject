package structs

import (
	"houseGoProject/enums/FamilyRoles"
	"houseGoProject/enums/Genders"
)

type HouseMember struct {
	Name    string
	Surname string
	Gender  Genders.Gender
	Age     int
	Role    FamilyRoles.FamilyRole
}
