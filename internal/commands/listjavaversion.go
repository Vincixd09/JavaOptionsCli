package commands

import (
	"fmt"

	u "github.com/Vincixd09/go-utils"
	"github.com/gookit/color"
)

func ListVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	color.Info.Prompt("Showing versions")

	color.Info.Prompt("Showing Java versions")
	err := u.RunCommandInteractive("update-alternatives", "--list", "java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	fmt.Print("\n")

	color.Info.Prompt("Showing javac versions")
	err = u.RunCommandInteractive("update-alternatives", "--list", "javac")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	fmt.Print("\n")
	color.Info.Tips("Press Enter to continue...")
	fmt.Scanln()

	fmt.Print("\033[3;J\033[H\033[2J")
}
