package main

import (
	"fmt"
)

//Extend program (5) to accept 3 inputs: “num1”, “num2” and “operation” where operation  could be “+”, “-”, “*” or “/” to  represent
// sum, difference, multiplication or division. Theoutput will be output of “num1” <operation> “num2”. The output shall be “num1=<num1> num2=<num2>
// <operation>=<result>” where “<operation>” will be replaced by the operation name. Use “sum”, “difference”, “multiply” and “divide” as an operation name

func main() {

	var (
		num1, num2 float64
		operator   string
	)
	fmt.Println("Enter num1: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Enter Proper Number:")
	}

	fmt.Println("Enter num2: ")
	_, err2 := fmt.Scan(&num2)
	if err2 != nil {
		fmt.Println("Enter Proper Number:")
	}

	fmt.Println("Enter operator : ")
	fmt.Scan(&operator)

	if operator != "+" && operator != "-" && operator != "/" && operator != "*" {

		fmt.Println("Enter Proper Arithmetic Operator ")
		return
	}

	result, lable, err := arithOperation(num1, num2, operator)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result is %f and operation is %s", result, lable)
		fmt.Println()
	}
}

// Function for performs Arithmetic Operation
func arithOperation(num1, num2 float64, operator string) (float64, string, error) {
	switch operator {
	case "+":
		return num1 + num2, "sum", nil
	case "-":
		return num1 - num2, "sbutraction", nil
	case "/":
		if num2 == 0 {
			fmt.Println(" Something wents wrong ")
			return 0, "Division", fmt.Errorf("Cannot Divide by zero!!")

		} else {
			return num1 / num2, "division", nil
		}

	case "*":
		return num1 * num2, "multiplication", nil

	default:
		return 0, "", fmt.Errorf("Something Wents Wrong!!")
	}

}
