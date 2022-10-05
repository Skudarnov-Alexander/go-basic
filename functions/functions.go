package main

import "fmt"

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func main() {

	fmt.Println(square)
	fmt.Println(circle)
	fmt.Println(triangle)

	myFigure := triangle

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	var x float64 = 2
	myArea := ar(x)

	fmt.Println(myArea)

}

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case 0:
		return func(side float64) (area float64) {
			return side * side
		}, true
	case 1:
		return func(side float64) (area float64) {
			return 3.14 * side * side
		}, true
	case 2:
		return func(side float64) (area float64) {
			return side * side / 2
		}, true
	default:
		return func(side float64) (area float64) {
			return 0
		}, false

	}

}
