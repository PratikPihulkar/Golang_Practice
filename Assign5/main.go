package main

import (
	"fmt"
)

// Write a program to prompt the user for 2 inputs: “num1” and “num2” and generate a sum of two numbers as output.
// The program must accept only whole numbers (positive or negative) or throw an error. The output shall be “num1=<num1> num2=<num2>
// sum=<result>” where “<num1>”, “<num2>” and “<result>” will be replaced with actual value.

func main() {

	var num1, num2 int

	fmt.Println("Enter First Number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Enter Proper Number!!")
		return
	}

	fmt.Println("Enter Second Number:")
	_, err2 := fmt.Scan(&num2)
	if err2 != nil {
		fmt.Println("Enter Proper Number!!")
		return
	}
	fmt.Println("Sum of two no is: ", num1+num2)

}
