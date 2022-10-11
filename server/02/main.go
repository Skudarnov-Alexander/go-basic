package main

import (
	"fmt"
	"io"
	"net/http"
)

// HelloWorld — обработчик запроса.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello, World</h1>"))
}

func GetHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
        return
	}
}

func QuaryHandler (w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("quary")
	if q == "" {
        http.Error(w, "The query parameter is missing", http.StatusBadRequest)
        return
    }
    // в нашем случае q примет значение "something"
    // продолжаем обработку запроса
    // ...
}

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
    // заголовки доступны методом Header.Get
    ct := r.Header.Get("Content-Type")
	_ = ct
    // для типового запроса ct примет значение "text/html; charset=UTF-8"
    // продолжаем обработку
    // ...
} 

func BodyHandler(w http.ResponseWriter, r *http.Request) {
    // читаем Body
	defer r.Body.Close()
    b, err := io.ReadAll(r.Body)
	_ = b
    // обрабатываем ошибку
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    // продолжаем обработку
    // ...
} 

func main() {
    // маршрутизация запросов обработчику
    http.HandleFunc("/", HelloWorld)
    // запуск сервера с адресом localhost, порт 8080
    err := http.ListenAndServe(":3333", nil)
    fmt.Println(err)
} 