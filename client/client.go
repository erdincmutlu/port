package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erdincmutlu/port/types"
)

func getPorts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := []types.Port{
		{

		},
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()

	router.
		HandleFunc("/portss", getPorts).
		Methods(http.MethodGet).
		Name("GetPorts")

	log.Fatal(http.ListenAndServe(":8080", router))
}
