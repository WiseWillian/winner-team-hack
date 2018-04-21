package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

type Location struct {
	Latitude string `json:"latitude"`
	Longintude string `json:"longitude"`
}

func Endpoint(w http.ResponseWriter, r *http.Request) {
	location := new (Location)

	location.Latitude = "111232.001"
	location.Longintude = "225546.102"

	jsonLocation, _ := json.Marshal(location)

	w.Write(jsonLocation)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/return", Endpoint).Methods("GET")

	return router
}

func main() {
	r := SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
