package main

import "fmt"

func main() {
	source := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len %d \n", len(source))
	fmt.Printf("cap %d \n", cap(source))
	s1 := append(source[:10], 99)

	fmt.Printf("len %d \n", len(s1))
	fmt.Printf("cap %d \n", cap(s1))
	s1[0] = 77
	fmt.Printf("source %v \n", source)
	fmt.Printf("s1 %v \n", s1)
	//s2 := append(s1, 99) 
	//fmt.Printf("len %d \n", len(s2))
	//fmt.Printf("cap %d \n", cap(s2))
	//fmt.Printf("v %v \n", source)
	//fmt.Printf("v %v \n", s2)
}