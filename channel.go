package main

import (
	"fmt"
	"runtime"
)

// channel as a parameter data type
func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	// learn go channel

	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	sayHelloTo := func(who string) {
		data := fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)

	// running channel as a data type parameter
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			data := fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
