package main

import (
	"fmt"
	"strings"
)

type studentMethod struct {
	name  string
	grade int
}

func (s studentMethod) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s studentMethod) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// method using pointer
func (s studentMethod) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to: ", name)
	s.name = name
}

func (s *studentMethod) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to: ", name)
	s.name = name
}

func main() {
	// method
	s1 := studentMethod{"john wick kurniawan", 21}
	s1.sayHello()

	name := s1.getNameAt(3)
	fmt.Println("sur name: ", name)

	// run method pointer
	fmt.Println("==Pointer on Method==")

	s3 := studentMethod{"john wick", 21}
	fmt.Println("s3 before", s3.name)

	s3.changeName1("jason bourne")
	fmt.Println("s3 after changeName1", s3.name)

	s3.changeName2("ethan hunt")
	fmt.Println("s3 after changeName2", s3.name)

	// access method from object variable
	s4 := studentMethod{"john wick", 22}
	s4.sayHello()

	// access method form object variable pointer
	s5 := &studentMethod{"ethan hunt", 23}
	s5.sayHello()

}
