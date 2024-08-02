package main

import (
	"fmt"
	"math"
)

type Cube struct {
	Side float64
}

func (c Cube) CubeVolume() float64 {
	return math.Pow(c.Side, 3)
}

func (c Cube) CubeArea() float64 {
	return math.Pow(c.Side, 2) * 6
}

func (c Cube) CubeCircumference() float64 {
	return c.Side * 12
}

func main() {
	var cCube Cube = Cube{4}
	fmt.Println(cCube.CubeVolume())
	fmt.Println(cCube.CubeArea())
	fmt.Println(cCube.CubeCircumference())
}
