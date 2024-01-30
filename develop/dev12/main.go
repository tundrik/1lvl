/*
*	12) Имеется последовательность строк - (cat, cat, dog, cat, tree) 
*	создать для нее собственное множество.
*
*************************************************/
package main

import (
	"fmt"
)


func main() {
	source := []string{"cat", "cat", "dog", "cat", "tree"}

	set := []string{}
	hash := map[string]struct{}{}


	for _, s := range source {
		if _, ok := hash[s]; !ok {
			// если на карте еще нет 
			// то добавляем и на карту
			// и в результат
			hash[s] = struct{}{}
			set = append(set, s)
		}
	}

	fmt.Printf("%v \n", set)
}