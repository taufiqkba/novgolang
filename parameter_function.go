package main

import (
	"fmt"
	"strings"
)

func main() {
	// function as a parameter
	var data = []string{"taufiq", "johne", "wick", "aguero", "henderson"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli \t\t: ", data)
	fmt.Println("Data contains O \t: ", dataContainsO)
	fmt.Println("Data length 5 \t\t: ", dataLenght5)

}

// using aliases closure scheme
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {

	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
