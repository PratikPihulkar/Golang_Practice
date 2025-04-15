package main

import (
	"fmt"
	"sort"
	"time"
)

// Same as (22) to accept a list of holidays, and then prompt a user for two inputs: input
// date as a first argument and a number of business days as a second argument. Number
// of business days will be a positive or negative whole number. The output shall be the
// date relative to an input date +/- the number of business days. Holidays must be
// excluded while calculating the output date. [Date manipulation & loop & boolean
// expressions, 4 hours]

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
func checkIsWorkingDay(holidayContainer []time.Time, dateStr time.Time) bool {

	if dateStr.Weekday().String() == "Saturday" || dateStr.Weekday().String() == "Sunday" {
		return false
	}
	for _, date := range holidayContainer {
		if date.Equal(dateStr) {
			return false
		}
	}
	return true
}

func findWorkingDaysCount(holidayContainer []time.Time, date1 time.Time, date2 time.Time) int {
	workingDayCount := 0
	for date := date1; !date.After(date2); date = date.AddDate(0, 0, 1) {
		if checkIsWorkingDay(holidayContainer, date) {
			workingDayCount++
		}
	}
	return workingDayCount
}

func main() {

	var holiday string
	var holidayContainer = []time.Time{}
	var date1 string
	var date2 string

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

	//Accept two dates for date range

EnterFirstDate:
	fmt.Println("Enter First Date :")
	fmt.Scan(&date1)
	validDate1, err := parseDate(date1)
	if err != nil {
		fmt.Println("date is in wrong format, Try Again:")
		goto EnterFirstDate
	}

EnterSecondDate:
	fmt.Println("Enter Second Date :")
	fmt.Scan(&date2)
	validDate2, err := parseDate(date2)
	if err != nil {
		fmt.Println("date is in wrong format, Try Again:")
		goto EnterSecondDate
	}

	//swap if daate not in order
	if validDate1.After(validDate2) {
		validDate1, validDate2 = validDate2, validDate1
	}

	workingDayCount := findWorkingDaysCount(holidayContainer, validDate1, validDate2)
	fmt.Println("workingDayCount : ", workingDayCount)

}
