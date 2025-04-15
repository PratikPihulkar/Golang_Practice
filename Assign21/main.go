package main

import (
	"fmt"
	"sort"
	"time"
)

/*. Write a program to accept a list of holidays (date in any of the supported formats). Store
this list in an internal array. After the user confirms entering of the holiday list, accept a
date from the user, and confirm whether it's a working day or not. (All Saturdays and
Sundays are implicitly considered as holidays). [Arrays & Date Manipulation, 4 hours]*/

// function for  convert string input to the date format
func parseDate(dateStr string) (time.Time, error) {

	dateFormats := []string{
		"02-01-2006",
		"2006-01-02",
		"01/02/2006",
		"January 2, 2006",
		"02 Jan 2006",
		"2 Jan 2006",
		"02-Jan-2006",
	}

	for _, format := range dateFormats {
		if tempararyUseDate, err := time.Parse(format, dateStr); err == nil {

			return tempararyUseDate, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", dateStr)
}

func checkIsRepeted(holidayContainer []time.Time, tempararyUseDate time.Time) bool {

	//date format is matched
	//so below for-loop for check date is present or not
	for _, singleDate := range holidayContainer {
		if singleDate.Equal(tempararyUseDate) {
			return true // it is repetes
		}
	}
	return false //not repeted
}

// function for check enterd date is loliday or not
func checkIsHoliday(holidayContainer []time.Time, dateStr time.Time) string {

	if dateStr.Weekday().String() == "Saturday" || dateStr.Weekday().String() == "Sunday" {
		return "It is holiday Because Of Week End !!"
	}
	for _, date := range holidayContainer {
		if date.Equal(dateStr) {
			return "is Holiday"
		}
	}
	return "is NOT Holiday"
}

func main() {

	var holiday string
	var holidayContainer = []time.Time{}
	var forCheckHoliday string
	fmt.Println("Enter Holidays and press 'done' for next Process ")
	fmt.Println("'saturday' and 'sunday' are already holidays ")
	for {

		fmt.Print(">>")
		fmt.Scan(&holiday)

		if holiday == "done" {
			// if Holiday Container is Empty
			if len(holidayContainer) == 0 {

				fmt.Println("No Holidays are Selected, so you cannot serch date !!")
				return
			} else {

				// Sort Holiday Container
				sort.Slice(holidayContainer, func(i, j int) bool {
					return holidayContainer[i].Before(holidayContainer[j])
				})
			}
			fmt.Println("Holidays fixed successfully")
			break
		}

		//Check date is in valid format or not
		//if Valid then add to  holiday container
		if newDate, err := parseDate(holiday); err == nil {
			if !checkIsRepeted(holidayContainer, newDate) {
				holidayContainer = append(holidayContainer, newDate)
			} else {
				fmt.Println("Date is Alredy Present!! ")
			}

		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Enter date for checking it is holiday or not :")
	fmt.Scan(&forCheckHoliday)

	var str string
	if dateForCheckInContainer, err := parseDate(forCheckHoliday); err == nil {
		fmt.Println(dateForCheckInContainer)
		str = checkIsHoliday(holidayContainer, dateForCheckInContainer)
	}

	fmt.Println(" Ans::  ", str)

}
