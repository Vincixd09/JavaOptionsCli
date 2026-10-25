package internal

import (
	"fmt"

	"github.com/gookit/color"
)

func ChangeVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	err := RunCommandInteractive("update-alternatives", "--config", "java")
	if err != nil {
		color.Error.Println("Error al intentar cambiar la version de java: ", err)
		return
	}

	err = RunCommandInteractive("update-alternatives", "--config", "javac")
	if err != nil {
		color.Error.Println("Error al intentar cambiar la version de javac: ", err)
		return
	}

	err = RunCommandInteractive("update-alternatives", "--config", "jar")
	if err != nil {
		color.Error.Println("Error al intentar cambiar la version de jar: ", err)
		return
	}

	fmt.Print("\033[3;J\033[H\033[2J")
}
