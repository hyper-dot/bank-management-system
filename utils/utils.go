package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetInput() string {
	var choice string
	fmt.Scan(&choice)
	return choice
}

// clears screen
func ClearScreen() {
	cmd := exec.Command("clear") // Use "cls" instead of "clear" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Returns to the Main Menu
func Return() {
	fmt.Println("Please Press 0 and Hit Enter to Return:")
	var choice string
	fmt.Scan(&choice)
	if choice == "0" {
		return
	}
	for choice != "0" {
		fmt.Println("Invalid Option !!")
		fmt.Println("Please Press 0 and Hit Enter to Return:")
		fmt.Scan(&choice)
	}
}

func Divider() {
	fmt.Println("---------------------------------------")
}

func DisplayMenu(menus []string) {
	ClearScreen()
	Divider()
	for _, item := range menus {
		fmt.Println(item)
	}
	Divider()
	fmt.Print("Plese select your an option: ")
}

func getChoice() {

}
