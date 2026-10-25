package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/pterm/pterm"
)

func NewJava() {

	fmt.Print("\033[3;J\033[H\033[2J")
	r, err := pterm.DefaultInteractiveTextInput.WithDefaultText("Ingresa la ruta del .tar de java").Show()
	if err != nil {
		pterm.Error.Printf("Error al ingresar el dato: %v\n", err)
		return
	}

	ExistDir()
	ExtracFile(r)

	maches := ExtractNameDir()
	Move(maches[0])

	pterm.Info.Println("Tar decomprimido correctamente")

	fmt.Println("Aquí tienes tus versiones de Java y sus carpetas")
	ViewOpt()

	pathR := "/opt/java/" + maches[0] + "/bin"

	pterm.Info.Println("Ej. Java 21 === 2100")
	version, err := pterm.DefaultInteractiveTextInput.WithDefaultText("Ingresa la version como en el ejemplo").Show()
	if err != nil {
		pterm.Error.Printf("Error al ingresar el dato: %v\n", err)
		return
	}

	java := pathR + "/java"
	javac := pathR + "/javac"
	jar := pathR + "/jar"
	archivos := []string{java, javac, jar}

	for _, archivo := range archivos {
		err := os.Chmod(archivo, 0755)
		if err != nil {
			pterm.Error.Printf("Error al dar permisos de ejecucion: %v\n", err)
			continue
		}
		pterm.Info.Printfln("Permisos de ejecución asignados a: %s\n", archivo)
	}

	err = RunCommandInteractive("update-alternatives", "--install", "/usr/bin/java", "java", java, version)
	if err != nil {
		pterm.Error.Println("Fallo al instalar java: ", err)
		return
	}

	err = RunCommandInteractive("update-alternatives", "--install", "/usr/bin/javac", "javac", javac, version)
	if err != nil {
		pterm.Error.Println("Fallo al instalar javac: ", err)
		return
	}

	err = RunCommandInteractive("update-alternatives", "--install", "/usr/bin/jar", "jar", jar, version)
	if err != nil {
		pterm.Error.Println("Fallo al instalar jar: ", err)
		return
	}
	fmt.Println("Java installed")

	rr, err := pterm.DefaultInteractiveConfirm.WithDefaultText(fmt.Sprintf("Quieres borrar el archivo %v?", r)).Show()
	if rr {
		err := os.Remove(r)
		if err != nil {
			pterm.Error.Println("Error al borrar el archivo: ", err)
			return
		}
	}

	time.Sleep(3 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
