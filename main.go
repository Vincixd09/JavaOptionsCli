package main

import (
	"JavaOptionsCli/internal"
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	fmt.Print("\033[3;J\033[H\033[2J")

	pterm.DefaultBox.
		WithBottomPadding(2).
		WithTopPadding(2).
		WithLeftPadding(10).
		WithRightPadding(10).
		Println(pterm.LightMagenta("JavaOptionsCli"))

	internal.Menu()
}
