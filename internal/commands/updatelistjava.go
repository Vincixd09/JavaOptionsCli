package commands

import (
	"JavaOptionsCli/internal/helps"
	"fmt"
	"time"

	"github.com/gookit/color"
)

func UpdateList() {
	fmt.Print("\033[3;J\033[H\033[2J")

	color.Info.Prompt("Updating system alternatives")

	err := helps.RunCommandInteractive("sudo", "update-alternatives", "--auto", "java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = helps.RunCommandInteractive("sudo", "update-alternatives", "--auto", "javac")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = helps.RunCommandInteractive("sudo", "update-alternatives", "--auto", "jar")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
