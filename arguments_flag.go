package main

import (
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
}
