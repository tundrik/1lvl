/*
*	10) Дана последовательность температурных колебаний:
*	-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
*	Объединить данные значения в группы с шагом в 10 градусов.
*	Последовательность в подмножноствах не важна.
*
*	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*
*************************************************/
package main

import (
	"fmt"
)

func main() {
	source := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := make(map[int][]float32)

	for _, t := range source {
		step := int(t/10) * 10
		result[step] = append(result[step], t)
	}

	fmt.Println(result)
}
