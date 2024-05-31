package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// ch for send data only
func sendMessage2(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

// ch for receive data only
func printMessage2(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

// implement timeout-channel on golang
func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func main() {
	// Range-Close on channel go
	fmt.Println("Range-Close on channel go")

	runtime.GOMAXPROCS(2)
	messages := make(chan string)
	go sendMessage2(messages)
	printMessage2(messages)

	// running implement timeout-channel
	fmt.Println("timeout-channel")
	messages2 := make(chan int)
	go sendData(messages2)
	retrieveData(messages2)
}
