package main

import (
	menu "bank/menus"
	"bank/utils"
)

func main() {
	for {
		utils.DisplayMenu(menu.MainMenu)
		choice := utils.GetInput()
		menu.HandleMainMenu(choice)
	}
}
