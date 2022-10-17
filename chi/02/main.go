package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var cars = map[string]string{
	"id1": "Renault",
	"id2": "BMW",
	"id3": "VW",
	"id4": "Audi",
}

func getAllCars(rw http.ResponseWriter, r *http.Request) {
	carList := carList()
	data := strings.Join(carList, ",")

	rw.Write([]byte(data))
}

func getCarBySlug(rw http.ResponseWriter, r *http.Request) {

}

func getCarByBrandAndModel(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	data := fmt.Sprintf("You are looking for: %s - %s", brand, model)
	rw.Write([]byte(data))

}

func getCarByBrand(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	data := fmt.Sprintf("You are looking for: %s", brand)
	rw.Write([]byte(data))

}

func addCar(rw http.ResponseWriter, r *http.Request) {

}

func deleteCar(rw http.ResponseWriter, r *http.Request) {

}

func getCarById(rw http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "carID")
	fmt.Printf("carID: %s\n", carID)
	if carID == "" {
		http.Error(rw, "carID in missed", http.StatusBadRequest)
		return
	}

	data, ok := carByID(carID)
	fmt.Printf("data: %s\n", data)

	if !ok {
		http.Error(rw, "carID is not exist", http.StatusBadRequest)
		return
	}

	rw.Write([]byte(data))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(timeTracer)

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Chi!!!"))
	})

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", getAllCars)
		r.Get("/{brand}", getCarByBrand)
		r.Get("/{brand}/{model}", getCarByBrandAndModel)
		r.Get("/{slug: [a-z-]+}", getCarBySlug)

		r.Route("/id/{carID}", func(r chi.Router) {
			r.Get("/", getCarById)
			r.Post("/", addCar)
			r.Delete("/", deleteCar)
		})
	})

	http.ListenAndServe(":8080", r)
}

func timeTracer (next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		handlerTime := time.Since(start)

		fmt.Println(handlerTime)
	})
}

func carList() []string {
	var list []string
	for _, v := range cars {
		list = append(list, v)
	}
	return list
}

func carByID(id string) (string, bool) {
	key := "id" + id
	if car, ok := cars[key]; ok {
		return car, ok
	}
	return "", false
}


