package main

import "fmt"

func main() {

	// looping for
	for i := 0; i < 5; i++ {
		fmt.Println("Number-i:", i)
	}

	// looping for just a condition
	var a = 0
	for a < 5 {
		fmt.Println("Number-a:", a)
		a++
	}

	// looping for without argument
	var b = 0
	for {
		fmt.Println("Number-b: ", b)
		b++

		if b == 5 {
			break
		}
	}

	// looping for range
	// string
	xs := "abc"
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}
	// array
	ys := [5]int{10, 20, 30, 40, 50}
	for _, v := range ys {
		fmt.Println("Value=", v)
	}
	// slice
	zs := ys[0:2]
	for _, v := range zs {
		fmt.Println("Value slice=", v)
	}
	// map
	kvs := map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for i, v := range kvs {
		fmt.Println("Key map=", i, "Value map=", v)
	}
	// ignore key and value
	for range kvs {
		fmt.Println("Done ignore key and value")
	}

	// looping with break and continue
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Number-i:", i)
	}

	// nested loop for
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// labels on looping
outerLoop:
	for i := 1; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Println("matrix[", i, "][", j, "]")
		}
	}
}
