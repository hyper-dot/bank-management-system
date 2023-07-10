package main

import (
	"bank/menus"
	"bank/utils"
)

func main() {
	for {
		menus.DisplayMainMenu()
		choice := utils.GetInput()
		menus.ProcessMainMenuChoice(choice)
	}
}
