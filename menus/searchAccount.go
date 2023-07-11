package menus

import (
	"bank/controller"
	"bank/utils"
	"fmt"
)

func SearchAccount() {
	for {
		utils.ClearScreen()
		fmt.Println("---------------------------------------")
		fmt.Println("1. Search By Phone")
		fmt.Println("2. Search by Account Number")
		fmt.Println("3. Search by Name")
		fmt.Println("0. Return")
		fmt.Println("---------------------------------------")
		fmt.Println()
		fmt.Print("Plese select your an option: ")

		choice := utils.GetInput()
		switch choice {
		case 1:
			controller.SearchByPhone()
		case 2:
			controller.SearchAccountbyAccNo()
		case 3:
			controller.SearchByName()
		case 0:
			return
		default:
			fmt.Println("invalid Options")
		}
	}
}
