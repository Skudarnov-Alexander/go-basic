package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var cars = map[string]string{
	"id1": "Renault",
	"id2": "BMW",
	"id3": "VW",
	"id4": "Audi",
}

func GetAllCars(rw http.ResponseWriter, r *http.Request) {
	carList := carList()
	data := strings.Join(carList, ",")

	rw.Write([]byte(data))
}

func GetCarById(rw http.ResponseWriter, r *http.Request) {
	carID := r.URL.Query().Get("id")

	if carID == "" {
		http.Error(rw, "carID in missed", http.StatusBadRequest)
		return
	}

	data := carByID(carID)
	fmt.Printf("data: %s\n", data)

	rw.Write([]byte(data))
}

func main() {
	http.HandleFunc("/cars", GetAllCars)
	http.HandleFunc("/car", GetCarById)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func carList() []string {
	var list []string
	for _, v := range cars {
		list = append(list, v)
	}
	return list
}

func carByID(id string) string {
	key := "id" + id
	if car, ok := cars[key]; ok {
		return car
	}
	return ""
}
