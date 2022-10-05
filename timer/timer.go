package main

import (
	"fmt"
	"time"
)

func main(){
	start := time.Now()
    timer := time.NewTimer(2 * time.Second) // создаём таймер
    t := <-timer.C                          // ожидаем срабатывания таймера
    fmt.Println(t.Sub(start).Seconds())     // выводим разницу во времени

	newTicker := time.NewTicker(2 * time.Second)

    for i:= 0; i < 10; i++ {
		c := <- newTicker.C
		fmt.Println(c.Sub(start).Seconds())
	}
}