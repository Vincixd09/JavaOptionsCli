package internal

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

func ExistDir() {
	info, err := os.Stat("/opt/java")
	if os.IsNotExist(err) {
		fmt.Println("Error: ", err)
		err := os.Mkdir("/opt/java", 0777)
		if err != nil {
			pterm.Error.Println(err)
		}
		return
	}

	if err != nil {
		pterm.Error.Println(err)
	}

	if info.IsDir() {
		pterm.Info.Println("Es un directorio y existe")
	} else {
		pterm.Info.Println("No es un directorio")
	}
}

func ExtracFile(tarName string) {
	file, err := os.Open(tarName)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	defer gzReader.Close()

	tr := tar.NewReader(gzReader)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			pterm.Error.Println(err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(header.Name, os.FileMode(header.Mode))
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				pterm.Error.Println(err)
			}
			defer outFile.Close()

			io.Copy(outFile, tr)
		}
	}
}

func Move(origen string) {
	destino := fmt.Sprintf("/opt/java/%v", origen)
	err := os.MkdirAll(destino, 0755)
	if err != nil {
		pterm.Error.Printf("Error al crear la carpeta: %v\n", err)
		return
	}

	err = filepath.Walk(origen, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			pterm.Error.Printf("Error al caminar por la carpeta: %v\n", err)
			return err
		}

		rel, err := filepath.Rel(origen, path)
		if err != nil {
			pterm.Error.Printf("Error al calcular la ruta relativa: %v\n", err)
			return err
		}

		rutaDestino := filepath.Join(destino, rel)

		if info.IsDir() {
			return os.MkdirAll(rutaDestino, info.Mode())
		}

		return copiarArchivo(path, rutaDestino, info.Mode())
	})
	if err != nil {
		pterm.Error.Printf("Error al mover la carpeta: %v\n", err)
		return
	}

	err = os.RemoveAll(origen)
	if err != nil {
		pterm.Error.Printf("Error al borrar la carpeta: %v\n", err)
		return
	}
}

func copiarArchivo(src, dst string, mode os.FileMode) error {
	origenFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer origenFile.Close()

	destinoFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer destinoFile.Close()

	_, err = io.Copy(destinoFile, origenFile)
	return err
}

func ExtractNameDir() []string {
	matches, err := filepath.Glob("jdk*")
	if err != nil {
		log.Fatal(err)
	}

	if len(matches) == 0 {
		log.Fatal("not found")
	}

	return matches
}

func RunCommandInteractive(command string, args ...string) error {
	fmt.Print("\n")
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunCommandWithOutput(command string, args ...string) error {
	fmt.Print("\n")
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ViewOpt() []string {
	dir, err := os.Getwd()
	if err != nil {
		pterm.Error.Printf("Error al conseguir la carpeta actual: %v\n", err)
		return []string{}
	}
	defer os.Chdir(dir)

	err = os.Chdir("/opt/java")
	if err != nil {
		pterm.Error.Printf("Error al moverse a /opt/java: %v\n", err)
		return []string{}
	}

	dirs, err := filepath.Glob("jdk*")
	if err != nil {
		pterm.Error.Printf("filepath.Glob fallo: %v\n", err)
		return []string{}
	}
	pterm.Info.Println(strings.Join(dirs, " | "))
	return dirs
}

func Menu() {
	fmt.Print("\033[3;J\033[H\033[2J")

	pterm.DefaultBox.
		WithBottomPadding(2).
		WithTopPadding(2).
		WithLeftPadding(10).
		WithRightPadding(10).
		Println(pterm.LightMagenta("JavaOptionsCli"))

	options := []string{
		"New Version Java",
		"Changes Version",
		"View List Versions",
		"Delete Version",
		"Exit",
	}

	action := map[string]func(){
		"New Version Java":   NewJava,
		"Changes Version":    ChangeVersion,
		"View List Versions": ListVersion,
		"Delete Version":     DeleteVersion,
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
