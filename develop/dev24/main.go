/*
*	24) Разработать программу нахождения расстояния между двумя точками,
*	которые представлены в виде структуры Point с инкапсулированными
*	параметрами x,y и конструктором.
*
*************************************************/
package main

import (
	"fmt"
	"math"
)


type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// возвращает расстояние между точками
func (p Point) distance(o Point) float64 {
	return math.Sqrt(math.Pow(p.x-o.x, 2) + math.Pow(p.y-o.y, 2))
}

func main() {
	p1 := New(64, 98)
	p2 := New(12, 45)
	fmt.Println(p1.distance(p2))
	fmt.Println(p2.distance(p1))
	fmt.Println(p1.distance(p1))
}
// 74.24957912338628
// 74.24957912338628
// 0