/*
*	3) Дана последовательность чисел: 2,4,6,8,10.
*	Найти сумму их квадратов
*	с использованием конкурентных вычислений.
*
*************************************************/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	source := [5]int64{2, 4, 6, 8, 10}

	var sum int64

	wg := sync.WaitGroup{}
	wg.Add(5)
	
	for i := range source {
		go func(i int) {
			// атомарно прибавляем 
			atomic.AddInt64(&sum, source[i] * source[i])
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("%d \n", sum)
}
