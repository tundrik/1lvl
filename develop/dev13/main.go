/*
*	13) Поменять местами два числа без создания временной переменной
*
*************************************************/
package main

import (
	"fmt"
)

func swap(a, b int) {
	// вроде тоже что и a, b = b, a
	fmt.Printf("a %d b %d\n", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a %d b %d\n", a, b)
}

func swap1(a, b int) {
	// будет проблема если результат сложения
	// не влезет в int
	fmt.Printf("a %d b %d\n", a, b)
	b = a + b
	a = b - a
	b = b - a
	fmt.Printf("a %d b %d\n", a, b)
}

func main() {
	swap(23, -45)
	swap1(23, -45)
}
