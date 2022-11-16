package main

import (
    "fmt"
    "os"
)

func main() {
    u := os.Getenv("USER")
    fmt.Println(u)

	envs := os.Environ()
	fmt.Println(envs)
} 