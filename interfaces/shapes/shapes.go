package main

import "fmt"

// shapes interface
type shape interface {
	getArea() float64
}

// print area of the shape using the type's getArea function
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// square struct with only one side length
type square struct {
	sideLength float64
}

// triangle struct with height and base lengths
type triangle struct {
	height float64
	base   float64
}

func (s *square) getArea() (area float64) {
	area = s.sideLength * s.sideLength
	return
}

func (t *triangle) getArea() (area float64) {
	area = 0.5 * t.base * t.height
	return
}
