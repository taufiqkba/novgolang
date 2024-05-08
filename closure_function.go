package main

import (
	"fmt"
)

func main() {
	// closure function
	fmt.Println("Closure Function")

	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 5, 6, 1, 2}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// Immediately-Invoked Function Expression (IIFE)
	fmt.Println("Imediately-Invoked Function Expression (IIFE)")
	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("Original number:", numbers2)
	fmt.Println("New number:", newNumbers)

	// closure as a return value (running findMax function)

	var numFindMax = 12
	var numbersFindMax = []int{12, 3, 6, 2, 1, 7, 8, 9, 5}
	var howMany, getNumbers = findMax(numbersFindMax, numFindMax)
	var theNumbers = getNumbers()

	fmt.Println("closure function as a return value")
	fmt.Println("numbers :", numbersFindMax)
	fmt.Printf("find \t: %d\n\n", numFindMax)

	fmt.Println("found\t: ", howMany)
	fmt.Println("value\t: ", theNumbers)
}

// closure as a return value
func findMax(numbers []int, max int) (int, func() []int) {
	var result []int
	for _, e := range numbers {
		if e <= max {
			result = append(result, e)
		}
	}

	var getNumber = func() []int {
		return result
	}

	return len(result), getNumber
}

func LanTitude(a, b, c int) {

}
