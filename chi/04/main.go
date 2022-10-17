package main

import (
	"encoding/json"
	"net/http"

	"github.com/Skudarnov-Alexander/go_basic/chi/04/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*chi.Mux
	Repo models.CarRepo
}

func (h *Handler) CarList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := h.Repo.GetAllCars()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		carsJson, err := json.Marshal(cars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(carsJson)
	}
}

func NewHandler(repo models.CarRepo) *Handler {
	h := &Handler{
		Mux:  chi.NewMux(),
		Repo: repo,
	}
	h.Get("/cars", h.CarList())
	return h
}

func main() {

}
