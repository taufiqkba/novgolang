package main

import (
	"fmt"
	"math/rand"
	"time"
)

// random string
var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(lenght int) string {
	b := make([]rune, lenght)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

// random
func main() {
	// randomizer := rand.New(rand.NewSource(10))
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random-1:", randomizer.Int())
	fmt.Println("random-2:", randomizer.Int())
	fmt.Println("random-3:", randomizer.Int())
	fmt.Println("random-4:", randomizer.Float32())
	fmt.Println("random-5:", randomizer.Uint32())
	fmt.Println("random-5:", randomizer.Intn(20)) // to show random number between n-1 (0-19)

	// running random string
	fmt.Println("random string with 5 character:", randomString(16))
}
