package main

import "fmt"

func main() {
	// заполняется значениями по умолчанию
	var myArrayA [5]int
	fmt.Println(myArrayA)

	myArrayB :=  [3]int{4,5,6}
	myArrayC :=  [5]int{4,5,6}
	fmt.Println(myArrayB)
	fmt.Println(myArrayC)

	myArrayD :=  [...]int{4,5,6,7,8}
	fmt.Println(myArrayD)
	fmt.Println(len(myArrayD))

	myArrayE :=  [7]int{2:3, 3:6}
	fmt.Println(myArrayE)

	myArrayF :=  [...]int{2:3, 3:6}
	fmt.Println(myArrayF)

	// Массив и цикл
	var weekTemp = [7]int{3, 5, -7, 9, 15, 5, 6}
	var avg float32
	var sumTemp int

	for _, v := range weekTemp {
		sumTemp += v
	}

	fmt.Printf("Суммарная температура - %d \n", sumTemp)

	avg = float32(sumTemp) / 7

	fmt.Printf("Средняя температура за неделю - %.02f \n", avg)

	

}