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
	r, err := pterm.DefaultInteractiveSelect.WithOptions(dirs).WithDefaultText("Elija la version de java que quiera borrar").Show()
	if err != nil {
		pterm.Error.Println("Error al seleccionar una opcion: ", err)
		return
	}
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

	rr, err := pterm.DefaultInteractiveConfirm.WithDefaultText("Quieres borrar los archivos de esta version de java, que esta ubicado en /opt/java/").Show()
	if err != nil {
		pterm.Error.Println("Error al borrar jar: ", err)
		return
	}

	if rr {
		err = RunCommandInteractive("sudo", "rm", "-rf", pathR)
		if err != nil {
			pterm.Error.Println("Error: ", err)
			return
		}
		pterm.Info.Println("Java System Deletion")
	} else {
		pterm.Info.Println("Java ya no figura en las alternativas del sistema, pero sigue estando en la carpeta /opt/java/")
	}
	time.Sleep(4 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
