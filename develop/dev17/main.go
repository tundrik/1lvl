/*
*	17) Реализовать бинарный поиск встроенными методами языка.
*
*************************************************/
package main

import (
	"fmt"
)

func binarySearch(source []int, target int) int {
	start := 0 // указатель на начало
	mid := 0
	end := len(source) - 1 // указатель на конец

	// пока не схлопнемся
	for start <= end {
		// вычесляем индекс середины
		mid = (start + end) >> 1
		switch {
		case source[mid] > target:
			end = mid - 1 // отбрасываем первую половину
		case source[mid] < target:
			start = mid + 1 // отбрасываем вторую половину
		default:
			return mid
		}
	}
	return -1
}

func main() {
	source := []int{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6}

	fmt.Printf("index %d \n", binarySearch(source, 6))
	fmt.Printf("index %d \n", binarySearch(source, -4))
	fmt.Printf("index %d \n", binarySearch(source, 61))
}
