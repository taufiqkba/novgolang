package main

import "fmt"

type person3 struct {
	name string
	age  int
}

func main() {
	// Anonymouse Struct
	fmt.Println("Anonymouse Struct")
	var s1 = struct {
		person3
		grade int
	}{}
	s1.person3 = person3{"wick", 22}
	s1.grade = 3

	fmt.Println("name: ", s1.person3.name)
	fmt.Println("age: ", s1.person3.age)
	fmt.Println("grade: ", s1.grade)
}
