package internal

import (
	"JavaOptionsCli/internal/commands"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func Menu() {
	options := []string{
		"New Version Java",
		"Changes Version",
		"View List Versions",
		"Delete Version",
		"Update List",
		"Exit",
	}

	action := map[string]func(){
		"New Version Java":   commands.NewJava,
		"Changes Version":    commands.ChangeVersion,
		"View List Versions": commands.ListVersion,
		"Delete Version":     commands.DeleteVersion,
		"Update List":        commands.UpdateList,
		"Exit": func() {
			fmt.Print("\033[3;J\033[H\033[2J")
			os.Exit(1)
		},
	}

	selected, _ := pterm.DefaultInteractiveSelect.WithOptions(options).WithMaxHeight(10).Show()
	if action, ok := action[selected]; ok {
		action()
		Menu()
	} else {
		pterm.Error.Println("Dont exits options")
	}

}
