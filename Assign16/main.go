package main

import (
	"fmt"
	"time"
)

/*
Write a program to print current date/time in following formats (one line per format) in
UTC time e.g. “16 Mar 2022” “Mar 16, 2022” “2022-03-16” “2022-03-16T15:52:00Z”
“Tuesday, 16 March 2022” [Date manipulation, 3 hours]

Extend (14) to print supported timezones, and accept a valid timezone as input and print
time as per the time zone selected by an end user. [Date manipulation & switch
statement, 3 hours]
*/

func main() {

	utcNow := time.Now()
	fmt.Println(utcNow.Format("02 Jan 2006"))
	fmt.Println(utcNow.Format("Jan 02, 2006"))
	fmt.Println(utcNow.Format("2006-01-02"))
	fmt.Println(utcNow.Format(time.RFC3339))
	fmt.Println(utcNow.Format("Monday, 02 January 2006"))

	fmt.Println()

	var location string

	fmt.Println("Choose one of them or type Proper TimeZone: \nAsia/Kolkata \nAmerica/New_York \nEurope/London \nAustralia/Sydney \nUTC")
	fmt.Scanln(&location)

	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Println("Enter Valid Location :")
		return
	}

	cTime := time.Now().In(loc)

	switch location {
	case "Asia/Kolkata":
		fmt.Println("Asia/Kolkata: ", cTime)
	case "America/New_York":
		fmt.Println("America/New_York", cTime)
	case "Europe/London":
		fmt.Println("Europe/London", cTime)
	case "Australia/Sydney":
		fmt.Println("Australia/Sydney", cTime)
	case "UTC":
		fmt.Println("UTC", cTime)

	}

}
