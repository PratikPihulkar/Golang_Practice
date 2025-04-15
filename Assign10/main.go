package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program to prompt for a sequence of numbers, one number at a time, till the user
enters “proceed”. Upon receiving “proceed”, the program shall calculate the sum of all
numbers and produce an output. Ensure that only valid numbers are considered as an
input; if you detect an invalid number input, give a meaningful error message and exit.
Your program must work correctly even if the user enters zero or just one number. [Data
types and validations & for-loop, 4 hours]

Extend (8) to support “sort” operation. Use an in-built function call for sorting numbers.
*/

func main() {
	fmt.Println("Enter Numbers ")
	numberContainer := []int{}
	newNumberContainer := []int{}
	var input string
	for {

		fmt.Print(">>")
		fmt.Scan(&input)

		if input == "sort" {
			if len(numberContainer) == 0 {
				fmt.Println("No numbers for Sorting ")
				return
			}
			break
		}

		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Enter Proper Number Only!!")
			continue
		} else {
			numberContainer = append(numberContainer, value)
		}
	}
	//calling function
	newNumberContainer = sortArr(numberContainer)
	fmt.Println(newNumberContainer)
}

// Function for Sorting an Array/SLice
func sortArr(numberContainer []int) []int {
	sort.Ints(numberContainer) //convert to string for sorting
	return numberContainer

}
