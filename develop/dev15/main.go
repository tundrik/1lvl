/*
*	15) К каким негативным последствиям может привести данный фрагмент кода,
*   и как это исправить? Приведите корректный пример реализации.
*
*   var justString string
*
*   func someFunc() {
*       v := createHugeString(1 << 10)
*       justString = v[:100]
*   }
*
*   func main() {
*       someFunc()
*   }
*
*************************************************/
package main

import (
	"fmt"
	"strings"
)


func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10)

	// justString = v[:100]
    // Утечка памяти, вызванная подстроками
	//
    // https://go101.org/article/memory-leaking.html
	//
    return strings.Clone(v[:100]) // allocation and copy, костыль конечно тот еще
}

func main() {
	justString := someFunc()
    fmt.Printf("cap %v\n", cap([]byte(justString)))
}
