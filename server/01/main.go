package main

import (
	"fmt"
	"net/http"
)

// HelloWorld — обработчик запроса.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello, World</h1>"))
}

func main() {
    // маршрутизация запросов обработчику
    http.HandleFunc("/", HelloWorld)
    // запуск сервера с адресом localhost, порт 8080
    err := http.ListenAndServe(":3333", nil)
    fmt.Println(err)
} 