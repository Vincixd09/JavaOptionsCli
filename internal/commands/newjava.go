package commands

import (
	"JavaOptionsCli/internal/helps"
	"fmt"
	"os"
	"time"

	u "github.com/Vincixd09/go-utils"
	"github.com/gookit/color"
)

func NewJava() {
	fmt.Print("\033[3;J\033[H\033[2J")

	var (
		tar     string
		path    string
		version string
	)
	fmt.Println("Enter the name of the tar file")
	fmt.Scanln(&tar)

	info, err := os.Stat("/opt/java")
	helps.ExistDir(info, err)

	helps.ExtracFile(tar)
	maches := helps.ExtractNameDir()
	pathF := "/opt/java/" + maches[0]
	err = u.RunCommandInteractive("sudo", "mv", maches[0], pathF)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}
	color.Info.Println("Tar decompress correctly")

	fmt.Println("Here are your Java versions and their folders")
	err = u.RunCommandWithOutput("lsd", "/opt/java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	color.Info.Tips("Ej. jdk-21")
	fmt.Println("Enter the folder for the new Java version")
	fmt.Scanln(&path)
	pathR := "/opt/java/" + path + "/bin"

	color.Info.Tips("Ej. Java 21 === 2100")
	fmt.Println("Enter the Java version")
	fmt.Scanln(&version)

	java := pathR + "/java"
	javac := pathR + "/javac"
	jar := pathR + "/jar"

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--install", "/usr/bin/java", "java", java, version)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--install", "/usr/bin/javac", "javac", javac, version)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--install", "/usr/bin/jar", "jar", jar, version)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}
	fmt.Println("Java installed")

	time.Sleep(3 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
