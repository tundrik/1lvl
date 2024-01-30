/*
*	7) Реализовать конкурентную запись данных в map.
*
*************************************************/
package main

import (
	"fmt"
	"sync"
)

type Store struct {
	mu sync.Mutex
	m  map[int]int
}

func New() Store {
	return Store{
		mu: sync.Mutex{},
		m:  make(map[int]int),
	}
}

func (s *Store) Set(k int, v int) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func main() {
	s := New()

	wg := sync.WaitGroup{}

	wg.Add(8)

	for i := 0; i < 8; i++ {
		go func(i int) {
			s.Set(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("s %v \n", s.m)
}

// s map[0:0 1:1 2:2 3:3 4:4 5:5 6:6 7:7]
