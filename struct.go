package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

type person struct {
	name string
	age  int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name: ", s1.name)
	fmt.Println("grade: ", s1.grade)

	// initiate struct object
	fmt.Println("==Initiate Struct Object==")
	var s2 = student{}
	s2.name = "wick"
	s2.grade = 2

	var s3 = student{"ethan", 2}
	var s4 = student{name: "jason", grade: 2}

	fmt.Println("student 2: ", s2.name)
	fmt.Println("student 3: ", s3.name)
	fmt.Println("student 4: ", s4.name)

	var s5 = student{grade: 5, name: "uzi"}
	fmt.Println("student 5: ", s5.name)

	// pointer object variable
	fmt.Println("==Pointer Object Variable==")
	var s6 = student{name: "taufiq", grade: 2}
	var s7 *student = &s6

	fmt.Println("student 1, name: ", s6.name)
	fmt.Println("student 6, name: ", s7.name)

	s6.name = "kurniawan"
	fmt.Println("student 1, name: ", s6.name)
	fmt.Println("student 6, name: ", s7.name)

	// combine slice & struct
	var allStudents = []person{
		{name: "John", age: 22},
		{name: "Wick", age: 23},
		{name: "Kurniawan", age: 24},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is ", student.age)
	}
}
