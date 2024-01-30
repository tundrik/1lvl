/*
*	14) Разработать программу, которая в рантайме
*	способна определить тип переменной:
*	int, string, bool, channel из переменной типа interface{}.
*
*************************************************/
package main

import (
	"fmt"
)

type human struct{}

func (c human) String() string { 
	return "I human" 
}

func whatTypeAmI(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("I int %d\n", v)
	case string:
		fmt.Printf("I string %s\n", v)
	case bool:
		fmt.Printf("I bool %v\n", v)
	case chan string:
		fmt.Printf("I chan string %v\n", v)
	case func():
		v()
	case human:
		fmt.Printf("%s\n", v)
	default:
		fmt.Printf("I don't know %v\n", v)
	}
}

func main() {
	whatTypeAmI(28487)
	whatTypeAmI("Go")
	whatTypeAmI(true)
	whatTypeAmI(make(chan string))
	whatTypeAmI(func() { fmt.Printf("I func\n") })
	whatTypeAmI(human{})
	whatTypeAmI(nil)
}

