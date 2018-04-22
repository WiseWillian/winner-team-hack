package main

import (
	"gopkg.in/zabawaba99/firego.v1"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Coordenada struct {
	Latitude string `json:"lat"`
	Longitude string `json:"lng"`
}

var f = firego.New("https://winner-hackthon.firebaseio.com/", nil)

func PostBip(w http.ResponseWriter, r *http.Request) {
	coordenada := new (Coordenada)

	body, err_reading := ioutil.ReadAll(r.Body)
	if(err_reading != nil) {
		fmt.Println("Erro ao ler dados da requisição!")
	}

	err_unmarshal := json.Unmarshal(body, &coordenada)
	if(err_unmarshal != nil) {
		fmt.Println("Erro ao transformar JSON em objeto")
	}

	pushedFirego, err := f.Push(coordenada)

	fmt.Printf("%s: %s\n", pushedFirego, err)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/post", PostBip).Methods("POST")
	return router
}

func main() {
	r := SetupRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
