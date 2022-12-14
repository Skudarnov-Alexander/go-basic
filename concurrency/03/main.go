package main

import (
	"fmt"
)

/*
Допишите программу так, чтобы горутина прочитала все числа из каналов a и b и отправила их в канал c.
Канал для чтения выбирается оператором select случайным образом.
Программа должна вывести числа от 10 до 29 в произвольном порядке.
*/

func thread(dec int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- (dec*10 + i)
		}
		close(ch)
	}()
	return ch
}

func main() {
	a := thread(1)
	b := thread(2)
	c := make(chan int)
	go func() {
		for a != nil || b != nil {
			select {
			case num, ok := <-a:
				if !ok {
					a = nil
					continue
				}

				c <- num
			case num, ok := <-b:
				if !ok {
					b = nil
					continue
				}

				c <- num

			}
		}
		close(c)

		// допишите код
		// добавьте цикл с оператором select
		// не забудьте в конце закрыть канал 'c'
		// ...
	}()
	for v := range c {
		fmt.Println(v)
	}
}
