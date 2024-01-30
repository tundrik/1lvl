/*
*	16) Реализовать быструю сортировку массива (quicksort)
*   встроенными методами языка.
*
*************************************************/
package main

import (
	"fmt"
)

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main() {
	source := []int{6, -5, 44, -3, -2, 90, 0, 1, 29, 3, -4, 5, -6}

	fmt.Printf("%v\n", quickSortStart(source))
}
