package main

import (
	"fmt"
	"strconv"
)

/*Extend program (5) to accept 3 inputs: “num1”, “num2” and “operation” where operationcould be “+”, “-”, “*” or “/” to
represent sum, difference, multiplication or division. Theoutput will be output of “num1” <operation> “num2”. The output shall
be “num1=<num1>num2=<num2> <operation>=<result>” where “<operation>” will be replaced by theoperation name. Use “sum”, “difference”,
“multiply” and “divide” as an operation namewhen printing the result. [If/then/else OR switch statement, 3 hours]

Extend (7) to accept statistical operation instead of “proceed”. Valid values for “<operation>” are count (to count number of valid numbers),
mean to calculate arithmetic mean, min to calculate minimum value (minimum of all numbers input), max for maximum value
(maximum of all numbers input). [Switch & functions, 4 hours]
*/

func main() {

	numberContainer := []int{}
	var nums string
	fmt.Println("Enter Numbers countinuously of print for operation:(also print for further operation 'mean', 'min', 'max')")
	var val int
	var err error

	for {
		fmt.Print(">>")
		fmt.Scan(&nums)

		if nums == "min" || nums == "max" || nums == "mean" {
			if len(numberContainer) == 0 {
				fmt.Println("Number Container is Empty!!")
				return
			}
			break
		}

		singleNumber, err := strconv.Atoi(nums)

		if err == nil {
			numberContainer = append(numberContainer, singleNumber)
			continue
		} else {
			fmt.Println("Enter Valid Number or Value!")
			continue
		}
	}

	switch nums {
	case "mean":
		val, err = findMeanOfNumbers(numberContainer)
	case "min":
		val, err = findMininumnNumber(numberContainer)
	case "max":
		val, err = findMaximumNumber(numberContainer)
	default:
		fmt.Println("Enter valid value")
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

}

func findMaximumNumber(numberContainer []int) (int, error) {

	maximumNumber := numberContainer[0]

	for _, value := range numberContainer {
		if value > maximumNumber {
			maximumNumber = value
		}
	}
	return maximumNumber, nil

}

func findMininumnNumber(numberContainer []int) (int, error) {
	minimumNumber := numberContainer[0]

	for _, value := range numberContainer {
		if value < minimumNumber {
			minimumNumber = value
		}
	}
	return minimumNumber, nil
}

func findMeanOfNumbers(numberContainer []int) (int, error) {
	sum := 0
	for _, val := range numberContainer {

		sum = sum + val

	}
	return sum / len(numberContainer), nil

}
