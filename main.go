package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Seed struct {
	ID     string  `json: "id"`
	Name   string  `json: "name"`
	Season *Season `json: "season"`
}

type Season struct {
	Type     string `json: "type"`
	Rainfall string `json: "rainfall"`
}

var seeds []Seed

func getSeeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seeds)
}

func deleteSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range seeds {
		if item.ID == params["id"] {
			seeds = append(seeds[:index], seeds[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(seeds)
}

func getSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range seeds {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var seed Seed
	_ = json.NewDecoder(r.Body).Decode(&seed)
	seed.ID = strconv.Itoa(len(seeds) + 1)
	seeds = append(seeds, seed)
	json.NewEncoder(w).Encode(seed)
}

func updateSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range seeds {
		if item.ID == params["id"] {
			seeds = append(seeds[:index], seeds[index+1:]...)
			var seed Seed
			_ = json.NewDecoder(r.Body).Decode(&seed)
			seed.ID = params["id"]
			seeds = append(seeds, seed)
			json.NewEncoder(w).Encode(seed)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	seeds = append(seeds, Seed{ID: "1", Name: "Sunflower", Season: &Season{Type: "Spring", Rainfall: "368mm"}})
	seeds = append(seeds, Seed{ID: "2", Name: "Wheat", Season: &Season{Type: "Autumn", Rainfall: "382mm"}})

	r.HandleFunc("/seeds", getSeeds).Methods("GET")
	r.HandleFunc("/seeds/{id}", getSeed).Methods("GET")
	r.HandleFunc("/seeds", createSeed).Methods("POST")
	r.HandleFunc("seeds/{id}", updateSeed).Methods("PUT")
	r.HandleFunc("seeds/{id}", deleteSeed).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
