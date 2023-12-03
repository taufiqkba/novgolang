package main

import "fmt"

func main() {
	var firstName string = "John"

	var lastName string = "Wick"
	fmt.Println("hallo", firstName, lastName+"!")
	fmt.Printf("hallo %s %s!\n", firstName, lastName)

	// variable without data type
	var fullName string = "John Wick"
	spareName := "John"

	fmt.Printf("My fullname is: %s,but my call is: %s \n", fullName, spareName)

	// multiple variable
	var first, second, third string
	// first, second, third = "one", "two", "three"
	one, isFriday, twoPointOne, say := 1, true, 2.1, "hello"
	fmt.Println(first, second, third)
	fmt.Println(one, isFriday, twoPointOne, say)

	// keyowrd new
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}
