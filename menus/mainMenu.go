package menus

import (
	"bank/utils"
	"fmt"
	"os"
)

func DisplayMainMenu() {
	utils.ClearScreen()
	fmt.Println("---------------------------------------")
	fmt.Println("1. Account Management")
	fmt.Println("2. Deposits and Withdrawals")
	fmt.Println("3. Payment Transfers")
	fmt.Println("4. Loan Processing")
	fmt.Println("5. Clearing and Settlement")
	fmt.Println("0. Exit")
	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Print("Plese select your an option: ")
}

func ProcessMainMenuChoice(choice int) {
	switch choice {
	case 1:
		ShowAccManagemntOpts()
	case 2:
		fmt.Println("You chose 2 !!")
	case 3:
		fmt.Println("You chose 3 !!")
	case 4:
		fmt.Println("You chose 4 !!")
	case 5:
		fmt.Println("You chose 5 !!")
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid Options")
	}
}
