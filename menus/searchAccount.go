package menu

import (
	"bank/utils"
	"fmt"
)

var searchAccountMenu = []string{
	"1. Search By Phone Number",
	"2. Search By Account Number",
	"3. Search By Name",
	"0. Return",
}

func SearchAccount() {
	utils.DisplayMenu(searchAccountMenu)
	for {
		choice := utils.GetInput()
		switch choice {
		case "1":
			SearchByPhone()
		case "2":
			SearchAccountbyAccNo()
		case "3":
			SearchByName()
		case "0":
			utils.DisplayMenu(AccountManagementMenu)
			return
		default:
			fmt.Println("Invalid Options")
		}
	}
}
