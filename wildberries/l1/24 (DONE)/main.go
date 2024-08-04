package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func MakePoint(x, y int) Point {
	return Point{x, y}
}

type Line struct {
	p1, p2 Point
}

func (l *Line) Length() float64 {
	return math.Sqrt(math.Abs((float64(l.p1.x) - float64(l.p2.x))) + math.Abs(float64(l.p1.y)-float64(l.p2.y)))
}

func main() {
	p1 := MakePoint(1, 1)
	p2 := MakePoint(0, 0)
	line := Line{p1, p2}
	fmt.Println(line.Length())
}
