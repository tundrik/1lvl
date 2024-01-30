/*
*	6) Реализовать все возможные способы остановки выполнения горутины.
*
********************************************************************
*
*	6-chan) Вариант с передачей канала в горутину и последующим его закрытием.
*   Подходит если писателей несколько
*
********/
package main

import (
	"fmt"
	"time"
)

func producer(stop <-chan struct{}, ch chan<- string) {
	for {
		select {
		case <-stop:
			fmt.Print("producer [stop]\n")
			return

		default:
			ch <- "countdown"
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func receiver(stop <-chan struct{}, ch <-chan string) {
	for {
		select {
		case msg := <-ch:
			fmt.Printf("msg: %v\n", msg)

		case <-stop:
			fmt.Print("receiver [stop]\n")
			return
		}
	}
}

func main() {
	stop := make(chan struct{})
	ch := make(chan string)

	go producer(stop, ch)
	go producer(stop, ch)

	go receiver(stop, ch)
	go receiver(stop, ch)

	time.Sleep(1000 * time.Millisecond)

	close(stop)

	time.Sleep(4 * time.Millisecond)
}

// msg: countdown
// ...
// receiver [stop]
// receiver [stop]
// producer [stop]
// producer [stop]
