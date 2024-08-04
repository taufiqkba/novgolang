package main

import (
	"fmt"
	"sync"
)

func doPrinting(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	fmt.Println("Wait Group")
	fmt.Println(" ")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrinting(&wg, data)
	}

	wg.Wait()
}
