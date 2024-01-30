/*
*	22) Разработать программу, которая
*	перемножает, делит, складывает,
*	вычитает две числовых переменных a,b,
*	значение которых > 2^20.
*
*************************************************/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("21000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("21000000000000000000000", 10)

	z := new(big.Int)
	fmt.Printf("%v * %v = %v\n", a, b, z.Mul(a, b))
	fmt.Printf("%v / %v = %v\n", a, b, z.Div(a, b))
	fmt.Printf("%v + %v = %v\n", a, b, z.Add(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, z.Sub(a, b))
}
