package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Encode-Decode Base64

	// Implement function EncodeToString() & DecodeToString()

	data := "john wick"
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodeString)

	decodeByte, _ := base64.StdEncoding.DecodeString(encodeString)
	decodeString := string(decodeByte)
	fmt.Println("decoded:", decodeString)

	// Encode & Decode function
	// encode
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	encodedString2 := string(encoded)
	fmt.Println(encodedString2)

	// decode
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	_, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodeString2 := string(decoded)
	fmt.Println(decodeString2)

	// Encode & Decode Data URL
	dataURL := "https://blibli.com/"

	encodedStringURL := base64.URLEncoding.EncodeToString([]byte(dataURL))
	fmt.Println(encodedStringURL)

	var decodeByteURL, _ = base64.URLEncoding.DecodeString(encodedStringURL)
	decodeStringURL := string(decodeByteURL)
	fmt.Println(decodeStringURL)

}
