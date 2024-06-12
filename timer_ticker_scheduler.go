package main

import (
	"fmt"
	"os"
	"time"
)

// Timer & Goroutine
func timerFunc(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntimeout! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func main() {
	fmt.Println("Time, Ticker, and Scheduler on Go")

	// time.Sleep
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 second")

	// scheduler using time.Sleep
	for true {
		fmt.Println("Hello!!")
		time.Sleep(1 * time.Second)
	}

	// time.NewTimer() function
	timer := time.NewTimer(4 * time.Second)
	fmt.Println("startTimer")
	<-timer.C
	fmt.Println("finishTimer")

	// time.AfteFunc()
	ch := make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	time.After()
	fmt.Println("time.After()")
	<-time.After(4 * time.Second)
	fmt.Println("expired")

	// Scheduler using Ticker
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second) // wait for 10 second
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello!!", t)
		}
	}

	// running timer & goroutine
	timeout := 5
	ch2 := make(chan bool)

	go timerFunc(timeout, ch2)
	go watcher(timeout, ch2)

	input := ""
	fmt.Print("what is 725/25?")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}
