package main

import (
	"fmt"
	"time"
)

// Extend (14) to print time in IST timezone. [Date manipulation, 2 hours]

func main() {

	fmt.Println("IST timezone:")

	ct := time.Now()
	fmt.Println("Time in IST format:", ct.Hour(), ":", ct.Minute(), ":", ct.Second())
	fmt.Println(ct.Weekday(), ",", ct.Day(), ct.Month(), ",", ct.Year())

}
