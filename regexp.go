package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Regexp -> Regular Expression
	fmt.Println("Regexp on Go")
	fmt.Printf("\n")

	// implementation regexp
	text := "banana burger and chicken soup"
	regexp, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// find on this text/string maximum 2 data
	res1 := regexp.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	//output: []string{"banan", "burger"}

	// find on this text/string all matches with regex
	res2 := regexp.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	//output: []string{"banan", "burger", "and", "chicken", "soup"}

	// method: regex.MatchString() -> to detect that string/text is valid with regex
	isMatch := regexp.MatchString(text)
	fmt.Println(isMatch) // output: true

	// method: regex.FindString() -> to find string that matches with regex criteria
	strFind := regexp.FindString(text)
	fmt.Println(strFind)

	// method: FindStringIndex() -> to return index from regex operation
	// this method is same with FindString(), only different the return, this method only return the index
	idx := regexp.FindStringIndex(text)
	fmt.Println(idx) //output: [0,6]

	strPrintIndex := text[0:6]
	fmt.Println(strPrintIndex) //output: "banana"

	// method: FindAllString() -> used to find many strings that meet the predefined regexp criteria
	// on the second parameter to define how many the return of strings, if -1 it will be return all of the strings
	strFindAll := regexp.FindAllString(text, -1)
	fmt.Println(strFindAll)

	strFindAll2 := regexp.FindAllString(text, 1)
	fmt.Println(strFindAll2)

	// method: ReplaceAllString() -> used to replace all of data string that matches with regex criteria with some conditional
	// this example will replace burger to potato
	strReplace := regexp.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})

	fmt.Println(strReplace)

	// method: Split() -> used to separate string that match with predefined regex criteria
	// regexp, err := regexp.Compile(`[a-b]+`)
	// split sparator is character "a" & "b"
	strSplit := regexp.Split(text, -1)
	fmt.Printf("%#v \n", strSplit)
	// []string{"", "n", "n", " ", "urger soup"}
}
