package main

import (
	"fmt"
	"runtime"
)

// simple goroutien example
func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// learn goroutie

	//running goroutine
	runtime.GOMAXPROCS(2)

	go print(5, "hello")
	print(5, "what's up")

	var input string
	fmt.Scanln(&input)
}
