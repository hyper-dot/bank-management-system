package menus

import (
	"bank/controller"
	"bank/utils"
	"fmt"
)

var DepositeAndWithdrawlMenu = []string{
	"1. Deposite Money",
	"2. Withdraw Money",
	"3. Transfer To An Account",
	"0. Return",
}

func ShowDepositeAndWithdrawlMenu() {
	for {
		utils.DisplayMenu(DepositeAndWithdrawlMenu)
		choice := utils.GetInput()
		switch choice {
		case "1":
			controller.Deposit()
		case "2":
		case "3":
			controller.TransferMoney()
		case "0":
			return
		default:
			fmt.Println("invalid Options")
		}
	}

}
