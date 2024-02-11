/*
*	9) Разработать конвейер чисел.
*	Даны два канала: в первый пишутся числа (x) из массива,
*	во второй — результат операции x*2,
*	после чего данные из второго канала должны выводиться в stdout.
*
********************************************************************/
package main

import (
	"fmt"
)

// начальный этап конвеера
func genesis(source []int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < len(source); i++ {
			ch <- source[i]
		}
		close(ch) // закрывает канал за собой
	}()
	return ch
}

// завершающий этап конвеера
func quadrate(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for n := range in {
			ch <- n * n
		}
		close(ch) // закрывает канал за собой
	}()
	return ch
}

func main() {
	source := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26}

	in := genesis(source)

	out := quadrate(in)

	for n := range out {
		fmt.Println(n)
	}
}

// https://go.dev/blog/pipelines
