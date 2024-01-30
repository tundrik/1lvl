/*
*	8) Дана переменная int64.
*	Разработать программу которая устанавливает i-й бит в 1 или 0.
*
*************************************************/
package main

import (
	"fmt"
)

// инвертирует указанный бит
func invertBit(n int64, pos uint) int64 {
	return n ^ (1 << pos) // XOR
}

func main() {
	var n int64 = 22

	fmt.Printf("\n%b \n", n)
	fmt.Printf("%d \n\n", n)

	fmt.Printf("%b \n", invertBit(n, 0))
	fmt.Printf("%d \n\n", invertBit(n, 0))

	fmt.Printf("%b \n", invertBit(n, 1))
	fmt.Printf("%d \n\n", invertBit(n, 1))

	fmt.Printf("%b \n", invertBit(n, 2))
	fmt.Printf("%d \n\n", invertBit(n, 2))
}

//https://lemire.me/blog/2023/02/07/bit-hacking-with-go-code/
