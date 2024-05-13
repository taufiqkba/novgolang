package main

import "fmt"

func main() {
	// Learn Pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value): ", numberA)
	fmt.Println("numberA (Address): ", &numberA)
	fmt.Println("numberB (value): ", *numberB)
	fmt.Println("numberB (address): ", numberB)

	// change value on pointer
	fmt.Println("")
	numberA = 5

	fmt.Println("numberA (value): ", numberA)
	fmt.Println("numberA (Address): ", &numberA)
	fmt.Println("numberB (value): ", *numberB)
	fmt.Println("numberB (address): ", numberB)

	// Pointer parameter
	var number = 4
	fmt.Println("before: ", number)

	change(&number, 10)
	fmt.Println("after: ", number)

}

func change(original *int, value int) {
	*original = value
}
