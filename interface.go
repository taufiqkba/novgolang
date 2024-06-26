package main

import (
	"fmt"
	"math"
	"strings"
)

type calculate interface {
	area() float64
	circumference() float64
}

// Circle
type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return c.diameter / 2
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circumference() float64 {
	return math.Pi * c.diameter
}

// Square
type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) circumference() float64 {
	return s.side * 4
}

// Embedded Interface
type calculate2D interface {
	area() float64
	circumference() float64
}

type calculate3D interface {
	volume() float64
}

type calculate2D3D interface {
	calculate2D
	calculate3D
}

type cube struct {
	side float64
}

func (c *cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func (c *cube) area() float64 {
	return math.Pow(c.side, 2) * 6
}

func (c *cube) circumference() float64 {
	return c.side * 12
}

// Casting an empty interface variable to a pointer object
type personS struct {
	name string
	age  int
}

func main() {
	var two_dimentional calculate

	two_dimentional = square{10.0}
	fmt.Println("==SQUARE==")
	fmt.Println("surface area:", two_dimentional.area())
	fmt.Println("circumference:", two_dimentional.circumference())

	two_dimentional = circle{14.0}
	fmt.Println("==CIRCLE==")
	fmt.Println("surface area:", two_dimentional.area())
	fmt.Println("circumference:", two_dimentional.circumference())
	fmt.Println("radius:", two_dimentional.(circle).radius())

	var two_dimentional_2 calculate = circle{14.0}
	var circle_shape circle = two_dimentional_2.(circle)

	fmt.Println(circle_shape.radius())

	// embedded interface
	var three_dimentional calculate2D3D = &cube{4}

	fmt.Println("====CUBE====")
	fmt.Println("Area:", three_dimentional.area())
	fmt.Println("Circumference:", three_dimentional.circumference())
	fmt.Println("Volume:", three_dimentional.volume())

	// Any or Empty Interface{}
	fmt.Println("==ANY or Empty Interface==")
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "pineapple", "banana"}
	fmt.Println(secret)

	secret = 12.5
	fmt.Println(secret)

	// another example case

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "pineapple", "banana"},
	}
	fmt.Println("using interface{} empty")
	fmt.Println(data)

	// any type data

	var dataAny map[string]any

	dataAny = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "pineapple", "banana"},
	}
	fmt.Println("using any")
	fmt.Println(dataAny)

	// casting variable any and empty interface
	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiple by 10 is:", number)

	secret = []string{"apple", "mango", "orange"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	// Casting an empty interface variable to a pointer object
	var secretPerson interface{} = &personS{name: "wick", age: 27}
	var name = secretPerson.(*personS).name
	fmt.Println(name)

	// Combined between Slice, Map, and Interface{}
	var personCombine = []map[string]interface{}{
		{
			"name": "wick",
			"age":  23,
		},
		{
			"name": "ethan",
			"age":  25,
		},
		{
			"name": "hunt",
			"age":  29,
		},
	}

	for _, each := range personCombine {
		fmt.Println(each["name"], "age is", each["age"])
	}

	// another example
	var fruitCombine = []interface{}{
		map[string]interface{}{
			"name":  "Strawberry",
			"total": 20,
		},
		[]string{
			"mango", "pineapple", "papaya", "orange",
		},
	}

	for _, fruit := range fruitCombine {
		fmt.Print(fruit)
	}
}
