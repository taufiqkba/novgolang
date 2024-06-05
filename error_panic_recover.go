package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// set custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// using recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error Ocured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	// A.37. Error, Panic & Recover
	fmt.Println("Error, Panic & Recover")
	var input string
	fmt.Print("Type some number:")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// example set custom error
	var name string
	fmt.Print("Type your name:")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		// fmt.Println(err.Error())
		panic(err.Error()) // using panic
		fmt.Println("end") // line after panic can't be execute
	}

	// implement recover
	defer catch()
	var nameRecover string
	fmt.Print("Type your name:")
	fmt.Scanln(&nameRecover)

	if valid, err := validate(nameRecover); valid {
		fmt.Println("Hello", nameRecover)
	} else {
		panic(err.Error())
		fmt.Println("end") // line after panic can't be execute
	}

	// implement recover on IIFE
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic occured", r)
	// 	} else {
	// 		fmt.Println("Application running perfectly")
	// 	}
	// }()
	// panic("some error happend")

	// another example recover on IIFE
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			// recover for IIFE on looping
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message: ", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			panic("some error happening")
		}()
	}
}
