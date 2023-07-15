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

// Returns to Menu
func ReturnTo(ParentMenu func()) {
	fmt.Printf("Please press 0 hit enter to return: ")
	choice := GetInput()

	for choice != "0" {
		fmt.Printf("Please press 0 hit enter to return: ")
		fmt.Scan(&choice)
	}
	ParentMenu()
	return
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
