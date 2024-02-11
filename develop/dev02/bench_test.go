package main

import (
	"fmt"
	"sync"
	"testing"
)


func BenchmarkOne(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*402)
	b.StartTimer()
	for i := len(source) - 1; i >= 0; i-- {
		source[i] = 74 * 6 * 9 / 6
	}
	b.StopTimer()
	fmt.Printf(" su %v \n", source[len(source)-10:])
}

func BenchmarkTwo(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*402)
	b.StartTimer()
	for i := len(source) - 1; i >= 0; i -= 2 {
		source[i] = 74 * 6 * 9 / 6
		source[i - 1] = 74 * 6 * 9 / 6
	}
	b.StopTimer()
	fmt.Printf(" su %v \n", source[len(source)-10:])
}


func BenchmarkMass(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*402)
	b.StartTimer()

	numCpu := 6

	part := len(source) / numCpu

	wg := sync.WaitGroup{}
	wg.Add(numCpu)
	b.StartTimer()
	for i := 0; i < numCpu; i++ {
		go func(i int) {
			offset := part * i // смещение откуда начинать
			for j := offset; j < offset+part; j++ {
				// мутируем елемент исходного масива
				source[j] = 74 * 6 * 9 / 6
			}
			wg.Done()
		}(i)
	}
	// ждем горутин
	wg.Wait()

	b.StopTimer()
	fmt.Printf(" as %v \n", source[len(source)-10:])
}
