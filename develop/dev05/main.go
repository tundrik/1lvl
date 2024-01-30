/*
*	5) Разработать программу,
*	которая будет последовательно отправлять значения в канал,
*	а с другой стороны канала — читать.
*	По истечению N секунд программа должна завершаться.
*
*************************************************/
package main

import (
	"fmt"
	"time"
)

// consumer выводит в stdout данные из канала
func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	// по истечению deadline программа завершается
	deadline := time.After(2 * time.Second)

	ch := make(chan string)

	go consumer(ch)

	for {
		select {
		case <-deadline:
			// завершаем
			return

		default:
			time.Sleep(200 * time.Millisecond)
			ch <- "Deadline is coming soon"
		}
	}

}
