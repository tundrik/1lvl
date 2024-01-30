/*
*	2) Написать программу, которая конкурентно рассчитает
*	значение квадратов чисел взятых из массива (2,4,6,8,10)
*	и выведет их квадраты в stdout.
*
************************************************************/
package main

import (
	"fmt"
	"runtime"
	"sync"
)



func main() {
	source := [...]uint32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26}
    
	numCpu := runtime.NumCPU() 

	// пробуем разделить исходную задачу
	// на разумное количество горутин
	//
	// если будет остаток посчитаем остаток 
	// в main 
	// смотрите ниже
    part := len(source) / numCpu

	wg := sync.WaitGroup{}

	for i := 0; i < numCpu; i++ {
		// запустим количество горутин равное NumCPU
		//
        // каждая горутина будет работать в рамках 
		// своего диапазона 

        // увеличим счетчик
		wg.Add(1)

		go func(i int) {
			// отложим уменьшение счетчика
            defer wg.Done()

			offset := part * i // смещение откуда начинать
			for j := offset; j < offset + part; j++ {
				// мутируем елемент исходного масива
				source[j] *= source[j] 

				fmt.Printf("G%d index %d\n", i, j)
			}
		}(i)
	}

    // если есть остаток то считаем
	for i := numCpu * part; i < len(source) ; i++ {
		// мутируем елемент исходного масива
		source[i] *= source[i] 

		fmt.Printf("Gmain index %d \n", i)
	}

    // ждем горутин
    wg.Wait()

	fmt.Printf("%v \n", source)
}
//Gmain index 12 
//G0 index 0
//G0 index 1
//G0 index 2
//G3 index 9
//G3 index 10
//G3 index 11
//G1 index 3
//G2 index 6
//G2 index 7
//G2 index 8
//G1 index 4
//G1 index 5
//[4 16 36 64 100 144 196 256 324 400 484 576 676] 