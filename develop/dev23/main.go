/*
*	23) Удалить i-ый элемент из слайса.
*
*************************************************/
package main

import (
	"errors"
	"fmt"
)

// возвращает новый слайс без указанного элемента
func remove(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) { // проверка границ
		return nil, errors.New("error bounds out of range")
	}
	ret := make([]int, 0, len(s)-1) // создаем новый слайс
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)
	return ret, nil
}

// возвращает урезанный слайс без указанного элемента
// без сохранения порядка
func fastRemove(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) { // проверка границ
		return nil, errors.New("error bounds out of range")
	}
	s[index] = s[len(s)-1]   // заменяем указанный на последний
	return s[:len(s)-1], nil // отрезаем последний
}

func main() {
	source := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(remove(source, 4))
	fmt.Println(source)

	fmt.Println(fastRemove(source, 4))
	fmt.Println(source) // UPS! [0 1 2 3 #9 5 6 7 8 #9]
}

//[0 1 2 3 5 6 7 8 9]
//[0 1 2 3 4 5 6 7 8 9]
//[0 1 2 3 9 5 6 7 8]
//[0 1 2 3 9 5 6 7 8 9]
