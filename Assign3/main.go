package main

import (
	"fmt"
	"os"
)

//Write a program to accept “<name>” as command line argument, and print “Hello <name>”. Replace “<name>” with the text entered by a user running the program

func main() {

	//First argument always name of file

	if len(os.Args) < 2 {
		fmt.Println("Enter the proper cmd line argument")
		return
	}
	fmt.Println("Hello ", os.Args[1])

}
