package main

import (
	"flag"
	"fmt"
)

func main() {
	MyIntFlagPointer := flag.Int("num", 5, "test decription")
	flag.Parse()
    fmt.Println("word:", *MyIntFlagPointer)
}
