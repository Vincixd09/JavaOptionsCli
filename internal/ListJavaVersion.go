package internal

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

func ListVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	pterm.Info.Println("Versiones de Java")
	err := RunCommandInteractive("update-alternatives", "--list", "java")
	if err != nil {
		pterm.Error.Println("Error al mostrar la lista de java: ", err)
		return
	}

	fmt.Print("\n")

	pterm.Info.Println("Versiones de Javac")
	err = RunCommandInteractive("update-alternatives", "--list", "javac")
	if err != nil {
		color.Error.Println("Error al mostrar la lista de javac: ", err)
		return
	}

	fmt.Print("\n")

	pterm.Info.Println("Versiones de Jar")
	err = RunCommandInteractive("update-alternatives", "--list", "jar")
	if err != nil {
		color.Error.Println("Error al mostrar la lista de jar: ", err)
		return
	}

	fmt.Print("\n")
	r, err := pterm.DefaultInteractiveConfirm.WithDefaultText("Quieres volver al inicio?").WithConfirmText("Yes").WithRejectText("No").Show()
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	if !r {
		ListVersion()
	}

	fmt.Print("\033[3;J\033[H\033[2J")
}
