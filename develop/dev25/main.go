/*
*	25) Реализовать собственную функцию sleep.
*
*************************************************/
package main

import (
	"fmt"
	"time"
)

func sleep(value time.Duration) {
	<-time.After(value)
}

func main() {
	fmt.Println("start sleep")
	sleep(1000 * time.Millisecond)
	fmt.Println("end sleep")
}
