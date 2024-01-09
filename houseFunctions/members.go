package houseFunctions

import (
	"context"
	"fmt"
	"houseGoProject/enums/FamilyRoles"
	"houseGoProject/enums/Genders"
	"houseGoProject/structs"
	"os"
)

func getHouseMembers(ctx context.Context, houseId int) []structs.HouseMember {
	conn := getConn(ctx)
	var members []structs.HouseMember
	membersRows, _ := conn.Query(ctx, fmt.Sprintf("%s%d", "SELECT * FROM housemember WHERE houseid=", houseId))

	for membersRows.Next() {
		var memberID, age, houseId int
		var name, surname string
		var gender Genders.Gender
		var role FamilyRoles.FamilyRole

		err := membersRows.Scan(&memberID, &name, &surname, &gender, &age, &role, &houseId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan member row data: %v\n", err)
			os.Exit(1)
		}

		member := structs.HouseMember{
			ID:      memberID,
			Gender:  gender,
			Name:    name,
			Age:     age,
			Role:    role,
			Surname: surname,
		}

		members = append(members, member)
	}

	return members
}
