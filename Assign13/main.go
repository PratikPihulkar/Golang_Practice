package main

import "fmt"

/*
Write a program to prompt for three inputs: character to be used for display, num1 to
represent number of rows and num2 to represent number of columns. The output will be
a rectangular matrix where each cell will print a character input as a first input value.
[Loops, 4 hours]
*/
func main() {

	var ch string
	var rows int
	var columns int

	fmt.Println("Enter the charector:")
	fmt.Scan(&ch)
	fmt.Println("Enter the rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the columns:")
	fmt.Scan(&columns)

	matrix(ch, rows, columns)
}
func matrix(ch string, r int, c int) {

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(ch, " ")
		}
		fmt.Println()
	}
}
