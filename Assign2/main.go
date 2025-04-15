package main

import "fmt"

// Write a program to prompt the user to enter a name, and print “Hello <name>”. Replace“<name>” with the text entered by the user
func main() {

	var name string

	fmt.Println("Enter Name:")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	// fmt.Printf("Hello %s \n",name)
}
