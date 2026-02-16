package helps

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

func ExistDir(info os.FileInfo, err error) {

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
		color.Info.Println("Is Dir and exist")
	} else {
		color.Info.Println("Exist but dont is a dir")
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

func ExtractNameDir() []string {
	matches, err := filepath.Glob("jdk-*")
	if err != nil {
		log.Fatal(err)
	}

	if len(matches) == 0 {
		log.Fatal("not found")
	}

	return matches
}
