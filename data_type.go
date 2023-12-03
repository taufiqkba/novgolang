package main

import "fmt"

func main() {
	// data type
	var positiveNumber uint8 = 89
	var negativeNumber = -12345

	fmt.Println("Positive Number:", positiveNumber)
	fmt.Printf("Negative Number: %d\n", negativeNumber)

	// data type decimal
	var decimalNumber = 2.89
	fmt.Printf("Decimal Number: %f\n", decimalNumber)
	fmt.Printf("Decimal Number: %.3f\n", decimalNumber)

	// data type boolean
	var exist bool = true
	fmt.Printf("is exist: %t\n", exist)

	// data type string
	var message string = "Hallo"
	fmt.Printf("this message is: %s\n", message)

	var anotherMessage = `My name is "John Wick".
	Hello, im there.
	Nice to meet you!`
	fmt.Println(anotherMessage)

	// data type nil & zero value
	// zero value is default value when data type declare without value
	var nilValue *int
	fmt.Println("Nil Value:", nilValue)

	var zeroValue int
	fmt.Println("Zero Value:", zeroValue)

	var zeroFloatValue float64
	fmt.Println("Zero Float Value:", zeroFloatValue)

	var zeroBoolValue bool
	fmt.Println("Zero Bool Value:", zeroBoolValue)

	var zeroStringValue string
	fmt.Println("Zero String Value:", zeroStringValue)
}
