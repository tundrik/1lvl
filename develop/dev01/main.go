/*
*   1)	Дана структура Human (с произвольным набором полей и методов).
*		Реализовать встраивание методов в структуре Action
*		от родительской структуры Human (аналог наследования).
*
***************************************************************/
package main

import "fmt"


type Human struct {
	name string
}

func (h Human) sitDown() {
	fmt.Printf("%s присел \n", h.name)
}

// Если у Action будут поля или методы - Human
// они затенят их
type Action struct {
	Human // встраивание

	count int
}

func (a *Action) train(s int) {
	for i := 0; i < s; i++ {
		a.count += 1
		// Вызов метода Human
		a.sitDown()
	}
}

func main() {
	action := new(Action)
	action.name = "Вася"

	action.train(5)
	fmt.Println(action.count)

	// Вызов метода Human
	action.sitDown()
    action.Human.sitDown()

	// Доступ к состаянию Human
	fmt.Println(action.name)
	fmt.Println(action.Human.name)
}
