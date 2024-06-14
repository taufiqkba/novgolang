package main

import (
	"fmt"
	"strings"
)

func main() {

	// strings.Contains() -> return bool
	isExists := strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	// strings.HasPrefix() -> to check string
	isPrefix1 := strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) //true

	isPrefix2 := strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) //false

	// strings.HasSuffix() -> endded with some strings
	isPrefix3 := strings.HasSuffix("john wick", "ic")
	fmt.Println(isPrefix3) //false

	isPrefix4 := strings.HasSuffix("john wick", "ck")
	fmt.Println(isPrefix4) //true

	// strings.Count() -> to calculate how many some character on string
	howMany := strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	// strings.Index() -> to found index on a string
	index1 := strings.Index("ethan hunt", "ha")
	fmt.Println(index1)
	index2 := strings.Index("ethan hunt", "n")
	fmt.Println(index2)

	// strings.Replace() -> to replace of string
	text := "banana"
	find := "a"
	replaceWith := "o"

	newText1 := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	newText3 := strings.Replace(text, find, replaceWith, -1) // to all substring
	fmt.Println(newText3)

	// strings.Repeat() -> to looping a string n times
	str := strings.Repeat("na", 5)
	fmt.Println(str)

	// strings.Split() -> to separate a string
	string1 := strings.Split("the dark night", " ")
	fmt.Println(string1) //output: ["the", "dark", "knight"]

	string2 := strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	// strings.Join() -> to combine string
	data := []string{"banana", "apple", "orange"}
	strJoin := strings.Join(data, "-")
	fmt.Println(strJoin) // banana-appple-orange

	// strings.ToLower() -> change string to lowercase
	strLower := strings.ToLower("aDdHuh")
	fmt.Println(strLower) // addhuh

	// strings.ToUpper() -> change string to UPPERCASE
	strUpper := strings.ToUpper("dhahar!")
	fmt.Println(strUpper) // DHAHAR!

}
