package main

import (
	"fmt"
	"strconv"
)

/*
Extend (7) to accept statistical operation instead of “proceed”. Valid values for
“<operation>” are count (to count number of valid numbers), mean to calculate
arithmetic mean, min to calculate minimum value (minimum of all numbers input), max
for maximum value (maximum of all numbers input). [Switch & functions, 4 hours]

Extend (9) to support “countodd” and “counteven” operations to respectively print
number of times odd numbers and number of even numbers found within the list.
[Expressions, 2 hours]
*/
func main() {

	fmt.Println("Enter Numbers")
	arr := []int{}

	var input string
	for {

		fmt.Print(">>")
		fmt.Scan(&input)

		val, err := strconv.Atoi(input)

		if err == nil {
			arr = append(arr, val)
			continue
		}
		if input == "countodd" || input == "counteven" {
			if input == "countodd" {
				_, odd := count(arr)
				fmt.Println("Odd no count is: ", odd)
			}
			if input == "counteven" {
				even, _ := count(arr)
				fmt.Println("Even no count is: ", even)
			}
			break
		} else {
			fmt.Println("Enetr Valid values!!")
		}

	}
}

func count(arr []int) (int, int) {

	countEven := 0
	countOdd := 0
	for _, val := range arr {

		if val%2 == 0 || val == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	return countEven, countOdd
}
