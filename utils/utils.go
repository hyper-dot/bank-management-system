package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetInput() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

// clears screen
func ClearScreen() {
	cmd := exec.Command("clear") // Use "cls" instead of "clear" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
