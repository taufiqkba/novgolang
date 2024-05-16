package main

import (
	"fmt"
	lib "novgolang/library"
)

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

	// running file on library
	// library.SayHello("ethan hunt")
	// library.introduce("ethan") // can't access because unexported function

	// var studentS1 = lib.Student{Name: "ethan", Grade: 2}
	// fmt.Println("name:", studentS1.Name)
	// fmt.Println("grade: ", studentS1.Grade)

	// imported file
	fmt.Printf("Name: %s\n", lib.Student.Name)
	fmt.Printf("Grade: %d\n", lib.Student.Grade)

}
