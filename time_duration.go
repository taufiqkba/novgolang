package main

import (
	"fmt"
	"time"
)

func main() {
	// Time Duration on Go
	fmt.Println("Time Duration")
	start := time.Now()
	time.Sleep(5 * time.Second)

	duration := time.Since(start)
	fmt.Println("time elapsed in seconds: ", duration.Seconds())
	fmt.Println("time elapsed in minutes: ", duration.Minutes())
	fmt.Println("time elapsed in hours: ", duration.Hours())

	// calculated between 2 time object
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration2 := t2.Sub(t1)
	fmt.Println("time elapsed in seconds: ", duration2.Seconds())
	fmt.Println("time elapsed in minutes: ", duration2.Minutes())
	fmt.Println("time elapsed in hours: ", duration2.Hours())

	// convertion number to time.Duration
	// example
	n := 5
	duration3 := time.Duration(n) * time.Second
	fmt.Println("duration is:", duration3)

}
