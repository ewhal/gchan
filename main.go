package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var DATABASE string
var configuration Configuration

type Configuration struct {
	// PORT for golang to listen on
	Port string
	// USERNAME database username
	Username string
	// PASS database password
	Password string
	// NAME database name
	Name string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func saveHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	DATABASE = configuration.Username + ":" + configuration.Password + "@/" + configuration.Name + "?charset=utf8"
	// create new mux router
	router := mux.NewRouter()

	// basic handlers
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/thread/{ID}", rootHandler)
	router.HandleFunc("/new", saveHandler)
	// ListenAndServe on PORT with router
	err = http.ListenAndServe(configuration.Port, router)
	if err != nil {
		log.Fatal(err)
	}

}
