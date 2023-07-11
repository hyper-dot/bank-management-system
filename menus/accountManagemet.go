package menus

import (
	accountController "bank/controller"
	"bank/utils"
	"fmt"
)

func ShowAccManagemntOpts() {
	for {
		utils.ClearScreen()
		fmt.Println("---------------------------------------")
		fmt.Println("1. Create Account")
		fmt.Println("2. Show All Accounts")
		fmt.Println("3. Search Account")
		fmt.Println("0. Return")
		fmt.Println("---------------------------------------")
		fmt.Println()
		fmt.Print("Plese select your an option: ")

		choice := utils.GetInput()
		switch choice {
		case 1:
			accountController.CreateAccount()
		case 2:
			accountController.ShowAllAccount()
		case 3:
			SearchAccount()
		case 0:
			return
		default:
			fmt.Println("invalid Options")
		}
	}

}
