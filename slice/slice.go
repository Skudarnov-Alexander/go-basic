package main

import "fmt"

func main() {
	var mySlice []int
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 5)
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", mySlice, len(mySlice), cap(mySlice))

	sliceA := make([]int, 0)
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", sliceA, len(sliceA), cap(sliceA))

	sliceB := make([]int, 5)
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", sliceB, len(sliceB), cap(sliceB))

	sliceC := make([]int, 5, 10)
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", sliceC, len(sliceC), cap(sliceC))

	fmt.Println("\n==Массив под слайсом==")

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}

	workDaysSlice := weekTempArr[:5]
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", workDaysSlice, len(workDaysSlice), cap(workDaysSlice))

	weekendSlice := weekTempArr[5:]
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", weekendSlice, len(weekendSlice), cap(weekendSlice))

	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice))

	weekTempSlice := weekTempArr[:]
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", weekTempSlice, len(weekTempSlice), cap(weekTempSlice))

	sliceD := workDaysSlice[:cap(workDaysSlice)]
	fmt.Printf("Слайс: %v, Длина: %d, Мощность: %d\n", sliceD, len(sliceD), cap(sliceD))

	fromTuesdayToThursDaySlice[2] = 0
	fmt.Println(fromTuesdayToThursDaySlice)
	fmt.Println(weekTempArr)
	fmt.Println(workDaysSlice)
	fmt.Println(weekTempSlice)

	fmt.Println("\n==Функуция append==")

	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a, len(a), cap(a))

	b := append(a, 6)
	fmt.Println(b, len(b), cap(b))

	b[1] = 666
	fmt.Println(a)
	fmt.Println(b)

	c := make([]int, 5, 7)
	fmt.Println(c)
	c = append(c, 1)
	fmt.Println(c)
	d := append(c, 2)
	d[1] = 2
	fmt.Println(c)
	fmt.Println(d)

	e := []int{0, 1, 2, 3, 4, 5, 6, 7}
	f := e[2:4:6]

	fmt.Println(e, len(e), cap(e))
	fmt.Println(f, len(f), cap(f))

	g := f[0:4]
	fmt.Println(g, len(g), cap(g))

	capSlice := []int{0, 1, 2, 3, 4, 5, 6}
	changeCapSlice(capSlice)

	fmt.Println("слайс исходный:", capSlice)

	fmt.Println("\n==Упражнение==")

	nums := make([]int, 100)
	for i:=0; i<100; i++ {
		nums[i] = i + 1
	}

	fmt.Println(nums)

	nums = append(nums[:10], nums[89:]...)
	fmt.Println(nums)
	for i:=0; i<10; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}

	fmt.Println(nums)


}

func changeCapSlice (s []int) {
	s = s[0:3:3]
	s[1] = 9
	fmt.Println("слайс внутри функции:",s)
	fmt.Println("cap слайса внутри функции:",cap(s))
}
