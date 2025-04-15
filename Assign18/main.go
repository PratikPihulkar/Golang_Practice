package main

import (
	"fmt"
	"time"
)

/*Write a program to accept a date and print whether the date falls in a leap year. Accept a
date in any format supported in one of the previous problems. [Date manipulation, 2 hours]*/

func main() {

	var date string
	fmt.Println("Enter the date:")
	fmt.Scan(&date)

	dt, err := time.Parse("02-01-2006", date)
	if err != nil {

		fmt.Println("Enter Date in Proper Format!!")
		return
	}
	year := dt.Year()

	fmt.Println(year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Entred year is leap year!")
	} else {
		fmt.Println("Entred year is NOT year!")
	}

}
