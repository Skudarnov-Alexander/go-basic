package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args
    fmt.Printf("All arguments: %v\n", args)
    // первый аргумент в списке — традиционно имя команды
    command := os.Args[0]
    fmt.Printf("Command name: %v\n", command)
    // далее слайс аргументов
    parameters := os.Args[1:]
    fmt.Printf("Parameters: %v\n", parameters)
}