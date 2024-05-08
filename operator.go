package main

import "fmt"

func main() {

	// aritmethic operator
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)

	// comparison operator
	var isEqual = (value == 2)

	fmt.Printf("Value %d (%t) \n", value, isEqual)

	// logical operator
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right: \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right: \t(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("left reverse: \t(%t) \n", leftReverse)

}
