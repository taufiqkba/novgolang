package main

import "fmt"

func main() {

	const firstName string = "John"
	fmt.Print("hallo", firstName, "!\n")

	const lastName = "Wick"
	fmt.Print("Hello, nice to meet you Mr. ", lastName, "!\n")

	// declare multiple constants
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	const (
		a = "Const"
		b // b have same value constant a
	)

	// another example
	const (
		today string = "Monday"
		now
		isToday2 = true
	)

	// constant in oneline
	const one, two, three = 1, 2, 3
	const four, five string = "four", "five"
}
