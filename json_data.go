package main

import (
	"encoding/json"
	"fmt"
)

// Decode JSON to Struct Object Variable

type UserJSONData struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// JSON Data
	var jsonString = `{"Name": "John Wick", "Age":27}`
	jsonData := []byte(jsonString)

	var data UserJSONData

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
