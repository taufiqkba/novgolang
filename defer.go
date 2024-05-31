package main

import (
	"fmt"
	"os"
)

func main() {
	// A.36. defer & exit on go
	fmt.Println("Defer & exit on go")
	// defer fmt.Println("hellooo!")
	// fmt.Println("welcome to the Island!")

	// orderSomeFood("pizza")
	// orderSomeFood("burger")

	// combine defer and IIFE
	number := 3
	if number == 3 {
		fmt.Println("hello 1")
		func() {
			defer fmt.Println("hello 3")
		}()
	}
	os.Exit(1) //implement os.Exit() -> it will be stop instantly the program
	fmt.Println("hello 2")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Please wait, thanks!")
	if menu == "pizza" {
		fmt.Println("That's a great choice!")
		fmt.Println("This pizza menu very delicious", "\n")
		return
	}

	fmt.Println("Your order is:", menu)
}
