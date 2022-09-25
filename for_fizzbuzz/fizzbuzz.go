package main

import "fmt"

func main() {
	for i := 0; i < 101; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}

	}
}
