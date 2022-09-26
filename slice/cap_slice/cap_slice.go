package main

import "fmt"

func main() {
	capSlice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("слайс исходный до функции:", capSlice)
	changeCapSlice(capSlice)

	fmt.Println("слайс исходный после функции:", capSlice)
	
}

func changeCapSlice (s []int) {
	s = s[0:3:3]
	s[1] = 9
	fmt.Println("слайс внутри функции:",s)
	fmt.Println("cap слайса внутри функции:",cap(s))
}