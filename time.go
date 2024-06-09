package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time on Go!")

	time1 := time.Now()
	fmt.Printf("time1 %v\n", time1)
	//2024-06-09 21:36:02.087933 +0700 WIB m=+0.000186079

	time2 := time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	// 2011-12-24 10:20:00 +0000 UTC

	now := time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())

	// parsing from string to time.Time
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

	// predifined layout format for parsing time
	var dateFormat, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(dateFormat.String()) // 2015-09-02 08:00:00 +0700 WIB

	// format time.Time to string
	var dateS1 = dateFormat.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1) // Wednesday 02, September 2015 08:00 WIB

	var dateS2 = dateFormat.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2) // 2015-09-02T08:00:00+07:00

	// error handling on parsing time.Time
	dateError, err := time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(dateError)

}
