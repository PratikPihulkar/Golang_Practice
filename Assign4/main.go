package main

import (
	"fmt"
	"os"
	"strings"
)

//Modify the above program to take two command-line parameters. The first is the string
//for the message template, like “Hello <name>” or any other template. The second is the
//actual name to print in the message by replacing “<name>” with the given name.

//args=== go run main.go "Hello {name}", "Pratik"

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Enter Proper Command line arguments")
		return
	}

	messege := os.Args[1]
	name := os.Args[2]

	newMessage := strings.ReplaceAll(messege, "{name}", name)
	// newMessage := strings.Replace(messege, "{name}", name, 1)

	fmt.Println(newMessage)

}
