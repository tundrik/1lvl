/*
*	6) Реализовать все возможные способы остановки выполнения горутины.
*
********************************************************************
*    
*	6-ctx) Вариант с передачей контекста в горутину и последующей его отменой.
*   Подходит если если нужна комбинация правил завершения
*   или контекст используется в иерархии горутин выше     
*
********/
package main

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context, ch chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("producer [cancel]\n")
			return

		default:
			ch <- "countdown"
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func receiver(ctx context.Context, ch <-chan string) {
	for {
		select {
		case msg := <-ch:
			fmt.Printf("msg: %v\n", msg)

		case <-ctx.Done():
			fmt.Print("receiver [cancel]\n")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	go producer(ctx, ch)
	go producer(ctx, ch)

	go receiver(ctx, ch)
	go receiver(ctx, ch)

	time.Sleep(1000 * time.Millisecond)

	cancel()

	time.Sleep(4 * time.Millisecond)
}

// msg: countdown
// ...
// receiver [cancel]
// receiver [cancel]
// producer [cancel]
// producer [cancel]
