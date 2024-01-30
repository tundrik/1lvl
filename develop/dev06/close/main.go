/*
*	6) Реализовать все возможные способы остановки выполнения горутины.
*
********************************************************************
*
*	6-close) Вариант с передачей канала в горутину и последующим его закрытием.
*	Подходит если писатель в канал один
*
********/
package main

import (
	"fmt"
	"time"
)

// завершает если канал закрыт
func receiver(ch <-chan string) {
	for {
		msg, ok := <-ch
		if !ok { // канал закрыт, msg nil
			fmt.Printf("ok: %v\n", ok)
			return
		}
		fmt.Printf("msg: %v, ok: %v\n", msg, ok)
	}
}

func main() {
	ch := make(chan string)

	go receiver(ch)

	ch <- "3 countdown"
	ch <- "2 countdown"
	ch <- "1 countdown"

	close(ch)

	time.Sleep(time.Millisecond)
}

// msg: 3 countdown, ok: true
// msg: 2 countdown, ok: true
// msg: 1 countdown, ok: true
// ok: false
