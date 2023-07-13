package menus

import (
	"fmt"
	"os"
)

var MainMenu = []string{
	"1. Account Management",
	"2. Deposits and Withdrawals",
	"3. Payment Transfers",
	"4. Loan Processing",
	"5. Clearing and Settlement",
	"0. Exit",
}

func ProcessMainMenuChoice(choice string) {
	switch choice {
	case "1":
		ShowAccManagemntOpts()
	case "2":
	case "3":
		fmt.Println("You chose 3 !!")
	case "4":
		fmt.Println("You chose 4 !!")
	case "5":
		fmt.Println("You chose 5 !!")
	case "0":
		os.Exit(0)
	default:
		fmt.Println("Invalid Options")
	}
}
