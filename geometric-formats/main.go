package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	newSquare := square{
		sideLength: 10,
	}
	printArea(newSquare)

	newTriangle := triangle{
		height: 10,
		base:   10,
	}
	printArea(newTriangle)
}
