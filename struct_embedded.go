package main

import "fmt"

type person2 struct {
	name string
	age  int
}

type student2 struct {
	grade int
	// age   int
	person2
}

func main() {
	// Embedded struct
	fmt.Println("==Embedded Struct==")

	var s1 = student2{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name is: ", s1.name)
	fmt.Println("age is: ", s1.age)
	fmt.Println("age is: ", s1.person2.age)
	fmt.Println("grade is: ", s1.grade)

	// Embedded struct with same property
	s1.person2.age = 22

	fmt.Println(s1.name)
	fmt.Println(s1.age)         // age of student
	fmt.Println(s1.person2.age) // age of person

	// Sub-Struct Value Filling
	fmt.Println("==Sub-struct value filling==")
	var p1 = person2{name: "wick", age: 32}
	var s2 = student2{person2: p1, grade: 2}

	fmt.Println("name: ", s2.name)
	fmt.Println("age: ", s2.age)
	fmt.Println("grade: ", s2.grade)

}
