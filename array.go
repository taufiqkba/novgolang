package main

import "fmt"

func main() {
	// array
	var names [4]string
	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Exco"
	names[3] = "Communicado"

	fmt.Println(names[0], names[1], names[2], names[3])

	// array over allocation
	var names2 [4]string

	names2[0] = "John"
	names2[1] = "Wick"
	names2[2] = "Exco"
	names2[3] = "Communicado"
	// names2[4] = "Error" // over allocation make error

	// array with initial value
	var names3 = [4]string{"John", "Wick", "Exco", "Communicado"}

	fmt.Println("Count element: ", len(names3))
	fmt.Println("Print element: ", names3)

	// array with vertical initial value
	var fruits [4]string

	/// horizontal
	fruits = [4]string{"Apple", "Orange", "Banana", "Grape"}

	/// vertical
	fruits = [4]string{
		"Apple",
		"Orange",
		"Banana",
		"Grape",
	}
	fmt.Println("Fuits name: ", fruits)

	// array without initial element
	numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Lenght array: ", len(numbers))
	fmt.Println("Print element: ", numbers)

	// array multidimension
	var numbers1 = [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	var numbers2 = [2][3]int{
		[3]int{2, 1, 2},
		[3]int{3, 4, 5},
	}

	fmt.Println(len(numbers1))
	fmt.Println(len(numbers2))

	// array with looping
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("=")

}
