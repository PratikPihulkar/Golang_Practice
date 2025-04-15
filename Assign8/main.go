package main

import (
	"fmt"
	"strconv"
)

/*Write a program to prompt for a sequence of numbers, one number at a time, till the user enters “proceed”. Upon receiving “proceed”,
the program shall calculate the sum of all numbers and produce an output. Ensure that only valid numbers are considered as an input;
if you detect an invalid number input, give a meaningful error message and exit. Your program must work correctly even if the user enters
zero or just one number. [Data types and validations & for-loop, 4 hours] */

func main() {

	var (
		sum     int
		input   string
		noCount int = 0
	)

	fmt.Println("Enter numbers or press 'proceed' to finish")

	for {
		fmt.Println("enter ", noCount+1, "th  Number >>")
		fmt.Scanln(&input)

		if input == "proceed" {
			break
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Enter valid Number Only! OR press 'proceed'")
			continue
		}
		noCount++
		sum = sum + num

	}
	if noCount == 0 {

		fmt.Println("No numbers Enterd")
		return

	} else {

		fmt.Println("Sum of Numbers is :", sum)
	}

}
