package main

import (
	"fmt"
	"strings"
)

/*Write a program to take the names of candidates as input. Prompt user to keep entering
new names till user enters “done”. Once a list of names are accepted, prompt the user
for a name. Output shall be “<name> exists” or “<name> does not exist”. A name exists
if the name exactly matches one of the names provided earlier. Use case insensitive
match for comparison. [Hash data structure & String operations, 6 hours]*/

func main() {

	//declear a map
	candidateContainer := make(map[string]bool)
	var input string
	fmt.Println("Enter names of candidate or enter 'done':")

	for {
		fmt.Print(":>> ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "done" {
			break
		}

		candidateContainer[input] = true
	}

	fmt.Println("Enter name for search in candidate list: ")
	fmt.Scanln(&input)
	if candidateContainer[input] {
		fmt.Println("Name is alredy Present")
	} else {
		fmt.Println("Entered name not present in list!")
	}

}
