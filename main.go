package main

import (
	"bank/menus"
	"bank/utils"
)

func main() {
	for {
		utils.DisplayMenu(menus.MainMenu)
		choice := utils.GetInput()
		menus.ProcessMainMenuChoice(choice)
	}
}
