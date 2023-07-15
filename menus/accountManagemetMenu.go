package menu

import (
	"bank/utils"
	"fmt"
)

var AccountManagementMenu = []string{
	"1. Create Account",
	"2. Show All Accounts",
	"3. Show Account Details",
	"4. Delete Account",
	"0. Return",
}

func ShowAccManagemntOpts() {
	utils.DisplayMenu(AccountManagementMenu)
	for {
		choice := utils.GetInput()
		switch choice {
		case "1":
			CreateAccount()
		case "2":
			ShowAllAccount()
		case "3":
			SearchAccount()
		case "4":
			DeleteAccount()
		case "0":
			return
		default:
			fmt.Printf("Please Enter Valid Option : ")
		}
	}
}
