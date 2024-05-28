package main

import (
	"fmt"
	"runtime"
)

func sendMessage2(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage2(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	// Range-Close on channel go
	fmt.Println("Range-Close on channel go")

	runtime.GOMAXPROCS(2)
	messages := make(chan string)
	go sendMessage2(messages)
	printMessage2(messages)
}
