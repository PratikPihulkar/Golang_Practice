package main

import (
	"fmt"
)

// Write a program to prompt the user for 2 inputs: “num1” and “num2” and generate a sum of two numbers as output.
//The program must accept only Floating Point numbers (positive or negative) or throw an error. The output shall be “num1=<num1> num2=<num2>
//sum=<result>” where “<num1>”, “<num2>” and “<result>” will be replaced with actual value.

func main() {
	var num1, num2 float64

	// fmt.Println("Enter Num1:")
	// fmt.Scan(&num1)
	// fmt.Println("Enter Num2:")
	// fmt.Scan(&num2)
	// fmt.Println("Sum: ", num1+num2)

	//Use Second Type
	fmt.Println("Enter First Number:")
	_, err := fmt.Scanf("%f", &num1)
	if err != nil {
		fmt.Println("Enter Proper Value!!")
		return
	}

	fmt.Println("Enter Second Number:")
	_, err2 := fmt.Scanf("%f", &num2)
	if err2 != nil {
		fmt.Println("Enter Proper Value!!")
		return
	}
	fmt.Println("Sum: ", num1+num2)

}
