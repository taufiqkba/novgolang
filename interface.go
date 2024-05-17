package main

import (
	"fmt"
	"math"
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

}
