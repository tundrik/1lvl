package main

import (
	"fmt"
	"sync"
	"testing"
	//"time"
)


func BenchmarkOne(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*360)
	b.StartTimer()
	for i := len(source) - 1; i >= 0; i-- {
		source[i] = 74 * 6 * 9 / 6
	}
	b.StopTimer()
	fmt.Printf(" su %v \n", source[len(source)-10:])
}

func BenchmarkTwo(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*360)
	b.StartTimer()
	for i := len(source) - 1; i >= 0; i -= 2 {
		source[i] = 74 * 6 * 9 / 6
		source[i - 1] = 74 * 6 * 9 / 6
	}
	b.StopTimer()
	fmt.Printf(" su %v \n", source[len(source)-10:])
}


func work(chank []uint16, wg *sync.WaitGroup) {
	for j := 0; j < len(chank); j++ {
		// мутируем елемент исходного масива
		chank[j] = 74 * 6 * 9 / 6
	}
	
	wg.Done()
}

func BenchmarkMass(b *testing.B) {
	b.StopTimer()
	source := make([]uint16, 1000000*360)
	b.StartTimer()

	numCpu := 4

	part := len(source) / numCpu

	wg := sync.WaitGroup{}
	wg.Add(numCpu)
	b.StartTimer()

	for i := 0; i < numCpu; i++ {
		offset := part * i // смещение откуда начинать
		chank := source[offset:offset+part]
		go work(chank, &wg)
	}

	// ждем горутин
	wg.Wait()

	b.StopTimer()
	fmt.Printf(" as %v \n", source[len(source)-10:])
}
