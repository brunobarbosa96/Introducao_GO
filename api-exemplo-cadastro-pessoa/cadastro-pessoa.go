package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/pessoa", getAll).Methods("GET")
	rotas.HandleFunc("/pessoa", create).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

type Pessoa struct {
	Name string
}

var pessoas = []Pessoa{

	Pessoa{Name: "Heisenberg"},
	Pessoa{Name: "Pinkman"},
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(pessoas)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var p Pessoa

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &p)

	pessoas = append(pessoas, p)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
