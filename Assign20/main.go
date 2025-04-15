package main

import (
	"fmt"
	"time"
)

/*
Write a program to accept two dates and print the count of week-end days (Consider
Saturdays and Sundays as week-ends). [Loops & Date manipulation & simple expressions, 6 hours]
*/

func main() {
	var date1, date2 string
	var totalDays int

	var holidayCount int
	var holidayCount2 int

	fmt.Println("Enter First Date: (DD-MM-YYYY)")

	fmt.Scan(&date1)
	newDate1, err := time.Parse("02-01-2006", date1)

	if err != nil {
		fmt.Println("Enter Date in Valid format!")
		return
	}

	fmt.Println("Enter Second Date: (DD-MM-YYYY)")

	fmt.Scan(&date2)
	newDate2, err := time.Parse("02-01-2006", date2)

	if err != nil {
		fmt.Println("Enter Date in Valid format!")
		return
	}

	//Date Swapping if order Not Matched
	if newDate1.After(newDate2) {
		temp := newDate1
		newDate1 = newDate2
		newDate2 = temp

		//short Way
		// newDate1, newDate2 = newDate2, newDate1
	}

	if newDate1.Equal(newDate2) {

		// dates are the same

		fmt.Println("Dates Are Same")
		if newDate1.Weekday().String() == "Saturday" || newDate1.Weekday().String() == "Sunday" {

			fmt.Println("Holiday Count: 1")
		} else {
			fmt.Println("Holiday Count: 0")
		}
		return
	}

	//Way 1:  For short date range:
	for d := newDate1; !d.After(newDate2); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {

			holidayCount++
		}
	}

	//way 2:  For Long date range:
	totalDays = int((newDate2.Sub(newDate1).Hours())/24) + 1

	//find weeks
	weeksInRange := int(totalDays / 7)

	// find extra days except week count (weeks + extra_days = date_range_count)
	extraDaysExceptWeek := int(totalDays % 7)

	usingDate := newDate1
	for i := 1; i <= extraDaysExceptWeek; i++ {

		if usingDate.Weekday() == time.Saturday || usingDate.Weekday() == time.Sunday {

			holidayCount2++
			usingDate.AddDate(0, 0, 1)

		} else {

			usingDate.AddDate(0, 0, 1)
		}
	}

	fmt.Println("TotalDays : ", totalDays)

	fmt.Println("Holiday count: (by using for loop)", holidayCount)

	fmt.Println("Holiday count: (by using shortcut way)", weeksInRange*2+holidayCount2)

}
