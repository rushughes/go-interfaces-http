package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	printArea() float64
}

func (s square) printArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) printArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	sq := square{sideLength: 10}
	tr := triangle{height: 10, base: 5}

	fmt.Println("Square area:", sq.printArea())
	fmt.Println("Triangle area:", tr.printArea())
}
