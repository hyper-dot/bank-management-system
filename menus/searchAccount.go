package menus

import (
	"bank/controller"
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
	for {
		utils.DisplayMenu(searchAccountMenu)
		choice := utils.GetInput()
		switch choice {
		case "1":
			controller.SearchByPhone()
		case "2":
			controller.SearchAccountbyAccNo()
		case "3":
			controller.SearchByName()
		case "0":
			return
		default:
			fmt.Println("invalid Options")
		}
	}
}
