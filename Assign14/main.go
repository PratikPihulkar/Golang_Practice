package main

import (
	"fmt"
	"time"
)

/*Write a program to print current date/time in following formats (one line per format) in
UTC time e.g. “16 Mar 2022” “Mar 16, 2022” “2022-03-16” “2022-03-16T15:52:00Z”
“Tuesday, 16 March 2022” [Date manipulation, 3 hours]*/

func main() {

	ct := time.Now().UTC()

	fmt.Println("Time in simple format: ", ct.Day(), ct.Month(), ",", ct.Year())
	fmt.Println()
	fmt.Println("Time in UTC format: ", ct)
	fmt.Println()
}
