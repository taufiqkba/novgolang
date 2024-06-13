package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Data type convertion on Go
	fmt.Println("Data type convertion on Go!")

	// strconv.Atoi() -> string to int
	str := "123"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	// strconv.Itoa() -> int to string
	num2 := 124
	str2 := strconv.Itoa(num2)
	fmt.Println(str2)

	// strconv.ParseInt() -> string to numerik 10 base
	str3 := "124"
	num3, err := strconv.ParseInt(str3, 10, 64) // (sourceNumber, biner bases, data type(int64))

	if err == nil {
		fmt.Println(num3)
	}
	// another example using bases 2 (biner)
	str4 := "1010"
	num4, err := strconv.ParseInt(str4, 2, 8)

	if err == nil {
		fmt.Println(num4)
	}

	// strconv.FormatInt() -> int64 to string
	num5 := int64(24)
	str5 := strconv.FormatInt(num5, 8)

	fmt.Println(str5) //30
	//octal representation of 24 is 30

	// strconv.ParseFloat() -> string to numerik
	str6 := "24.12"
	num6, err := strconv.ParseFloat(str6, 32)

	if err == nil {
		fmt.Println(num6) // 24.1200008392334
	}

	// strconv.FormatFloat()
	num7 := float64(24.12)
	str7 := strconv.FormatFloat(num7, 'f', 6, 64)
	fmt.Println(str7)

	// strconv.ParseBool() -> string to bool
	str8 := "true"
	bul, err := strconv.ParseBool(str8)

	if err == nil {
		fmt.Println(bul) //true
	}

	// strconv.FormatBool()
	bul2 := true
	str9 := strconv.FormatBool(bul2)

	fmt.Println(str9)

	// data type convertion using casting technique
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)

	// casting string <-> byte
	text1 := "hello"
	var byte = []byte(text1)

	fmt.Printf("%d %d %d %d\n", byte[0], byte[1], byte[2], byte[3])

	// byte2 := []byte{123, 1233, 134, 123}
	// s := string(byte2)

	// fmt.Printf("%s\n", s)

	var c int64 = int64('h')
	fmt.Println(c) //104

	var d string = string(104)
	fmt.Println(d) //h

	// assertions on any or empty interface

	data := map[string]interface{}{
		"name":    "john wick",
		"grade":   2,
		"height":  165.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	// casting to know data type on interface{} using switch
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
