package mcdeploy

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func NewQuestion(question string, action func(questionoutput string) string) string {

	var questionoutput string

	fmt.Println(question)

	fmt.Print(">> ")

	fmt.Scanf("%s", &questionoutput)

	return action(questionoutput)

}

func GetFileNameFromURL(url string) string {

	// Split the URL by "/"

	tokens := strings.Split(url, "/")

	// This gets us the filename of "craftbukkit.jar"

	file_name := tokens[len(tokens)-1]

	return file_name

}

func CreateFile(filename string, contents string) {

	_, err := os.Create(filename)

	// To write to files you have to use a array of bytes not a string sd we convert it

	ioutil.WriteFile(filename, []byte(contents), 0755)

	perror(err)

}

func GenerateServer() {

	servername := NewQuestion("Enter server name (this will generate a new folder)", func(questionoutput string) string {

		// Make a folder with the server name and give it read, write, & execute permissions

		os.Mkdir(questionoutput, 0777)

		// Change directory into this new folder (Note: The program does not move itself )

		os.Chdir(questionoutput)

		url := "https://dl.bukkit.org/latest-rb/craftbukkit.jar"

		// This gets the filename fromthe url and downloads it

		GetDownload(url, GetFileNameFromURL(url))

		return questionoutput

	})

	NewQuestion("Generate startup commands for your OS? (Note: It will not be run) (y or n)", func(questionoutput string) string {

		if strings.ToLower(questionoutput) == "yes" || strings.ToLower(questionoutput) == "y" {

			switch GetOS() {

			case "linux":

				filename := "start.sh"

				CreateFile(filename, Linuxscript)

			case "mac":

				filename := "start.command"

				CreateFile(filename, Macscript)

				// Give the file correct permissions to be executed

				err := exec.Command("chmod", "a+x", filename).Start()

				perror(err)

			case "windows":

				filename := "start.bat"

				CreateFile(filename, Windowsscript)

			}

		} else {

			fmt.Printf("New server created with name of %s", servername)

		}

		return "" // Required

	})

}

func NewBukkitServer() {

	GenerateServer()

}
