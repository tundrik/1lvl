/*
*	21 Реализовать паттерн «адаптер» на любом примере.
*
*************************************************/
package main

import (
	"fmt"
)

// импортируем
// нет возможности рефакторить
type WBLog struct {}

func (w WBLog) SendLog(s string) {
    fmt.Printf("WB %s\n", s)
}

// то что требуется
// например для сторонней библиотеки
type Logger interface {
    Log(s string)
}

// то что имеем
type WBLogger interface {
	SendLog(s string) 
}

// адаптер
type LoggerAdapter struct {
    WBLog
}

func (a LoggerAdapter) Log(s string) {
    a.SendLog(s)
}


func main() {
	var logger Logger = LoggerAdapter{WBLog{}}

    logger.Log("log из сторонней библиотеки")
}