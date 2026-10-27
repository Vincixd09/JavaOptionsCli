package internal

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

func DeleteVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	pterm.Info.Println("Borrar version de Java")

	dirs := ViewOpt()
	dirs = append(dirs, "Back")
	r, err := pterm.DefaultInteractiveSelect.WithOptions(dirs).WithDefaultText("Elija la version de java que quiera borrar").Show()
	if err != nil {
		pterm.Error.Println("Error al seleccionar una opcion: ", err)
		return
	}

	if r != "Back" {

		rr, err := pterm.DefaultInteractiveConfirm.WithDefaultText(fmt.Sprintf("Quieres desinstalar java de las update-alternative? ten encuenta que al confirmar tambien se van a borrar los archivos que estan ubicados en /opt/java/%v", r)).Show()
		if err != nil {
			pterm.Error.Println("Error al borrar jar: ", err)
			return
		}
		if rr {
			pathR := "/opt/java/" + r

			java := pathR + "/bin/java"
			javac := pathR + "/bin/javac"
			jar := pathR + "/bin/jar"

			err = RunCommandInteractive("sudo", "update-alternatives", "--remove", "java", java)
			if err != nil {
				pterm.Error.Println("Error al borrar java: ", err)
				return
			}

			err = RunCommandInteractive("sudo", "update-alternatives", "--remove", "javac", javac)
			if err != nil {
				pterm.Error.Println("Error al borrar javac: ", err)
				return
			}

			err = RunCommandInteractive("sudo", "update-alternatives", "--remove", "jar", jar)
			if err != nil {
				pterm.Error.Println("Error al borrar jar: ", err)
				return
			}

			err = RunCommandInteractive("sudo", "rm", "-rf", pathR)
			if err != nil {
				pterm.Error.Println("Error: ", err)
				return
			}
			pterm.Info.Println("Java System Deletion")
		} else {
			pterm.Info.Println("Java no sera borrado, volviendo al menu...")
		}
		time.Sleep(4 * time.Second)
		fmt.Print("\033[3;J\033[H\033[2J")
	} else {
		Menu()
	}
}
