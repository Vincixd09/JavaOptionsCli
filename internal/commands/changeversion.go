package commands

import (
	"JavaOptionsCli/internal/helps"
	"fmt"

	"github.com/gookit/color"
)

func ChangeVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	color.Info.Prompt("Showing the options")

	err := helps.RunCommandInteractive("sudo", "update-alternatives", "--config", "java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = helps.RunCommandInteractive("sudo", "update-alternatives", "--config", "javac")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = helps.RunCommandInteractive("sudo", "update-alternatives", "--config", "jar")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	fmt.Print("\033[3;J\033[H\033[2J")
}
