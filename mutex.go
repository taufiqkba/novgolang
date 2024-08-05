package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	// sync.Mutex
	val int
}

func (c *counter) Add(int) {
	// c.Lock()
	c.val++
	// c.Unlock()
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	fmt.Println("Mutex on Go!")

	// example case
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	var mtx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				meter.Add(1)
				mtx.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
