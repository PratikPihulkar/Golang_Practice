package main

import (
	"fmt"
	"time"
)

/*
Write a program to accept two dates (any of the supported period) and print an output
whether date1 and date2 are equal, date1 is earlier than date2 or date1 is later than
date2. [Date comparison, 1 hours]
*/

func main() {

	var date1, date2 string
	fmt.Println("Enter First Date: (DD-MM-YYYY)")

	fmt.Scan(&date1)
	dt1, err := time.Parse("02-01-2006", date1)

	if err != nil {
		fmt.Println("Enter Date in Valid format!")
		return
	}

	fmt.Println("Enter Second Date: (DD-MM-YYYY)")

	fmt.Scan(&date2)
	dt2, err := time.Parse("02-01-2006", date2)

	if err != nil {
		fmt.Println("Enter Date in Valid format!")
		return
	}

	if dt1.Before(dt2) {
		fmt.Println("First date come before second date:")
	} else if dt1.After(dt2) {
		fmt.Println("First date come after second date:")
	} else {
		fmt.Println("Dates Are Same")
	}

}
