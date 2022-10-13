package main

import (
	"log"
	"net/http"
	"github.com/Skudarnov-Alexander/go_basic/test/04/handlers"
)

func main() {
	http.HandleFunc("/status", handlers.StatusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}