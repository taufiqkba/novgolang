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
	// Decode JSON to Struct Object Variable
	var jsonString = `{"Name": "John Wick", "Age":27}`
	jsonData := []byte(jsonString)

	var data UserJSONData

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user : ", data.FullName)
	fmt.Println("age :", data.Age)

	// Decode JSON to map[string]interface{} & interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age : ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	decodeData := data2.(map[string]interface{})
	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age :", decodeData["Age"])

	// Decode JSON Array to Obeject Array

	jsonString2 := `[
		{"Name": "John Wick", "Age": 27},
		{"Name": "Ethan Hunt", "Age": 30}
	]`

	var data3 []UserJSONData

	err2 := json.Unmarshal([]byte(jsonString2), &data3)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	// Encode Object to JSON String
	object := []UserJSONData{{"john wick", 27}, {"ethan hunt", 32}}
	jsonData2, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonString2 = string(jsonData2)
	fmt.Println(jsonString2)

}
