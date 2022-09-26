package main

import (
	"fmt"
	
)

func main() {
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var sum int
	products := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for k, v := range products{
		if v > 500{
			fmt.Printf("%s - %d\n",k, v)
		}
	}

	for _, v := range order{
		sum += products[v]
	}

	fmt.Println(sum)

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	} 

	myMap := make(map[string]int)
	
	
	var idx int

	for _, v := range input{
		if _,ok := myMap[v]; !ok {
			myMap[v] = idx
			idx++
		}
	}

	
	output := make([]string, len(myMap))

	for k, v:= range myMap{
		output[v] = k
	}

	fmt.Println(myMap)
	fmt.Println(output)
}
