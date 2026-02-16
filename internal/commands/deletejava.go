package commands

import (
	"fmt"
	"time"

	u "github.com/Vincixd09/go-utils"
	"github.com/gookit/color"
)

func DeleteVersion() {
	fmt.Print("\033[3;J\033[H\033[2J")

	var (
		path string
		op   string
	)

	color.Info.Prompt("Delete Java version")

	err := u.RunCommandWithOutput("lsd", "/opt/java")
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}
	color.Info.Tips("Ej. jdk-21")
	fmt.Scanln(&path)
	pathR := "/opt/java/" + path + "/bin"

	java := pathR + "/java"
	javac := pathR + "/javac"

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--remove", "java", java)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	err = u.RunCommandInteractive("sudo", "update-alternatives", "--remove", "javac", javac)
	if err != nil {
		color.Error.Println("Error: ", err)
		return
	}

	color.Info.Prompt("Do you want to delete the Java files? in /opt/java")
	color.Info.Tips("yes or not")
	fmt.Scanln(&op)

	if op == "yes" {
		color.Info.Prompt("Enter the Java version you want to delete")
		color.Info.Tips("Ej. jdk-21")
		fmt.Scanln(&path)
		pathR = "/opt/java/" + path
		err = u.RunCommandInteractive("sudo", "rm", "-rf", pathR)
		if err != nil {
			color.Error.Println("Error: ", err)
			return
		}
		color.Info.Prompt("Java System Deletion")
	}
	color.Info.Prompt("Java is no longer in the system alternatives, but it's still in the folder /opt/java/ ")
	time.Sleep(4 * time.Second)
	fmt.Print("\033[3;J\033[H\033[2J")
}
