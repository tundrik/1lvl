/*
*	18) Реализовать структуру-счетчик,
*	которая будет инкрементироваться в конкурентной среде.
*	По завершению программа должна выводить итоговое значение счетчика.
*
*************************************************/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Statistic структура для хранения 
// количества событий
type Statistic struct {
	event atomic.Uint64
}

func New() Statistic {
	return Statistic{
		event: atomic.Uint64{},
	}
}

// Event атомарно увеличевает 
// количество событий на 1
func (s *Statistic) Event() {
	s.event.Add(1)
}

// State возвращает количество событий
// на текущий момент
func (s *Statistic) State() uint64 {
	return s.event.Load()
}

func main() {
	stc := New()

	wg := sync.WaitGroup{}

	wg.Add(666)

	for i := 0; i < 666; i++ {
		go func(i int) {
			stc.Event()
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("Statistic event: %d \n", stc.State())
}

//Statistic event: 666
