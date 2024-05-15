package main

import (
	"fmt"
)

type student struct {
	name  string `tag1` //using tag can used for encode/decode data especially on JSON
	grade int    `tag2`
}

type person struct {
	name string
	age  int
}

// implement nested struct
type studentNested struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

// declare and initiate struct with horizontal model
// type persons struct {name string; age int; hobbies []string}

// type struct alias
type Car struct {
	color string
	power int
}

type BrandCar = Car

type People1 struct {
	name string
	age  int
}

type People2 = struct {
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

	// declare and initiate struct with horizontal model
	// var p1 = struct{name string; age int}{age: 22, name: "wick"}
	// var p2 = struct{name string; age int}{age: 23, name: "john"}

	// type aliases
	fmt.Println("==Struct Aliases==")
	var car1 = Car{"Toyota", 2000}
	fmt.Println(car1)

	var car2 = BrandCar{"Honda", 3000}
	fmt.Println(car2)

	car := Car{"Suzuki", 2500}
	fmt.Println(BrandCar(car))

	brandCar := BrandCar{"Mazda", 3500}
	fmt.Println(Car(brandCar))

	// another cases
	type Number = int
	var num Number = 15
	fmt.Println(num)
}
