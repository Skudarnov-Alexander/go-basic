package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(1991, 11, 5, 19, 58, 0, 0, time.Now().Location())

	switch year := date.Year(); {
	case year > 2012:
		fmt.Println("hi alpha")
	case year > 1996:
		fmt.Println("hi zoomer")
	case year > 1980:
		fmt.Println("hi millenial")
	case year > 1964:
		fmt.Println("hi X")
	case year > 1945:
		fmt.Println("hi boomer")
	default:
		fmt.Println("hi smb")
	}
}
