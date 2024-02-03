package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	a := make([]int8, 1, 3)
	fmt.Println(len(a))
	fmt.Println(cap(a))

    a[0] = 66

	fmt.Println(len(a))
	fmt.Println(cap(a))

	someAction(a, 6)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println(a)
}
