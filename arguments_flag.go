package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Arguments & Flag

	// Using arguments
	argsRaw := os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	args := argsRaw[1:]
	fmt.Printf("-> %#v\n", args)

	// Using flag
	name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 25, "type you age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)

	dataName := flag.String("name", "anonymous", "type your name")
	fmt.Println(*dataName)

	// flag declare with pass reference using data store

	data1 := flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	data2 := flag.String("gender", "male", "type your gender")
	fmt.Println(*data2)

}
