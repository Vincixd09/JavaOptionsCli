package commands

import (
	"fmt"
	"time"

	u "github.com/Vincixd09/go-utils"
	"github.com/gookit/color"
)

func UpdateList() {
	fmt.Print("\033[3;J\033[H\033[2J")

	color.Info.Prompt("Updating system alternatives")

	err := u.RunCommandInteractive("sudo", "update-alternatives", "--auto", "java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--auto", "javac")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
