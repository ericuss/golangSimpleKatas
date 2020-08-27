package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	beer "github.com/ericuss/golangSimpleKatas/pkg"

	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository beer.BeerRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo beer.BeerRepository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/beers", a.fetch).Methods(http.MethodGet)
	r.HandleFunc("/beers/{Id:[a-zA-Z0-9_]+}", a.fetchById).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetch(w http.ResponseWriter, r *http.Request) {
	beers, _ := a.repository.Fetch()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}

func (a *api) fetchById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		_ = fmt.Errorf("Beer not found")
	}
	beer, err := a.repository.FetchById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Beer Not found")
		return
	}

	json.NewEncoder(w).Encode(beer)
}
