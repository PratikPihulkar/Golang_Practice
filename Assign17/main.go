package main

import (
	"fmt"
	"time"
)

/*17. Write a program to accept two dates (any of the formats supported in the earlier
problem) and print a difference in human readable format e.g. “1 year 2 day 32 minutes”.
[Date manipulation, 3 hours]*/

func main() {

	var date1, date2 string
	var layout = "02-01-2006"

	fmt.Println("Enter First Date:(Format: DD-MM-YYYY)")
	fmt.Scan(&date1)
	dt1, err := time.Parse(layout, date1)

	if err != nil {
		fmt.Println("Enter date in Proper Format")
		return
	}

	fmt.Println("Enter First Date:(Format: DD-MM-YYYY)")
	fmt.Scan(&date2)
	dt2, err := time.Parse(layout, date2)
	if err != nil {
		fmt.Println("Enter date in Proper Format")
		return
	}

	var difference int

	if dt1.After(dt2) {
		difference = int((dt2.Sub(dt1).Hours() / 24))
	} else if dt1.Before(dt2) {
		fmt.Println(" small")
		difference = int((dt1.Sub(dt2)).Hours() / 24)
	} else {
		fmt.Println("Dates Are Same")
	}

	if difference < 0 {
		difference = -difference // Make sure it's always positive
	}

	fmt.Println("The Difference Between two dates is: ", difference)

}
