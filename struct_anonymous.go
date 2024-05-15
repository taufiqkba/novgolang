package main

import "fmt"

type person3 struct {
	name string
	age  int
}

func main() {
	// Anonymous Struct
	fmt.Println("Anonymouse Struct")

	// anonymous struct witout filling property
	var s1 = struct {
		person3
		grade int
	}{}
	s1.person3 = person3{"wick", 22}
	s1.grade = 3

	fmt.Println("name: ", s1.person3.name)
	fmt.Println("age: ", s1.person3.age)
	fmt.Println("grade: ", s1.grade)

	// anonymous with property filling
	var s2 = struct {
		person3
		grade int
	}{
		person3: person3{"john", 25},
		grade:   5,
	}

	fmt.Println("name: ", s2.person3.name)
	fmt.Println("age: ", s2.age)
	fmt.Println("grade: ", s2.grade)

	// initiate slice for anonymous struct
	var allStudents = []struct {
		person3
		grade int
	}{
		{person3: person3{"john", 23}, grade: 4},
		{person3: person3{"wick", 24}, grade: 5},
		{person3: person3{"bayu", 25}, grade: 6},
	}

	for _, student := range allStudents {
		fmt.Println("name is: ", student.name, "age: ", student.age, "on grade: ", student.grade)
	}

	// Declaration of Anonymous Struct Using Keyword var

	var studentAnonymous struct {
		person3
		grade int
	}

	// another type version
	var studentAnonymous2 = struct {
		grade int
	}{
		15,
	}

	studentAnonymous.person3 = person3{"garuda", 22}
	studentAnonymous.grade = 1

	fmt.Println(studentAnonymous)
	fmt.Println(studentAnonymous2)

}
