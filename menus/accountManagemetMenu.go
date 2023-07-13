package menus

import (
	"bank/controller"
	"bank/utils"
	"fmt"
)

var AccountManagementMenu = []string{
	"1. Create Account",
	"2. Show All Accounts",
	"3. Show Account Details",
	"0. Return",
}

func ShowAccManagemntOpts() {
	utils.DisplayMenu(AccountManagementMenu)
	for {
		choice := utils.GetInput()
		switch choice {
		case "1":
			controller.CreateAccount()
		case "2":
			controller.ShowAllAccount()
		case "3":
			SearchAccount()
		case "0":
			return
		default:
			fmt.Println("Invalid Options!! Please Select a valid option: ")
		}
	}

}
