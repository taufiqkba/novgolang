package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	// Hash SHA1 -> Secure Hash Algorithm 1
	// algorithm to encrypt data one wa

	text := "this is secret words"
	sha := sha1.New()
	sha.Write([]byte(text))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)

	// running Salting method on Hash SHA1
	fmt.Printf("original text: %s\n", text)

	hashed1, salt1 := doHashUsingSalt(text)
	fmt.Printf("hashed1: %s\n", hashed1)
	// b9105e4da347937b1307aac78288acc71f146714

	hashed2, salt2 := doHashUsingSalt(text)
	fmt.Printf("hashed2: %s\n", hashed2)
	// ee6528bbcf60daaa07ee3a2b73c28e16a7b3543b

	hashed3, salt3 := doHashUsingSalt(text)
	fmt.Printf("hashed3: %s\n", hashed3)
	// 7d3d88276b7fa4bd2088e7c94b58cf9262bc14e2

	_, _, _ = salt1, salt2, salt3

}

// Salting method on Hash SHA1
func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt:'%s'", text, salt)
	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
