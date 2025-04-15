package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Write a program to take the names of candidates as input. Prompt user to keep entering
new names till user enters “done”. Once a list of names are accepted, prompt the user
for a search pattern (regex). Output shall be a list of all names where the search pattern
exists. [Array & Loop & Regex, 8-12 hours]
*/

func checkPatternMatchedValue(candidateContainer map[string]bool, pattern string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return
	}

	fmt.Println("\nNames matching the pattern:")
	found := false
	for name := range candidateContainer {
		if regex.MatchString(name) {
			fmt.Println("->", name)
			found = true
		}
	}

	if !found {
		fmt.Println("No matches found.")
	}
}

func main() {

	//declear a map
	candidateContainer := make(map[string]bool)
	var input string
	var pattern string

	fmt.Println("Enter names of candidate or enter 'done':")

	for {
		fmt.Print(":>> ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "done" {
			break
		}
		if input != "" {
			candidateContainer[input] = true
		}
	}

	fmt.Println("Enter search pattern: ")
	fmt.Scanln(&pattern)

	checkPatternMatchedValue(candidateContainer, input)
}
