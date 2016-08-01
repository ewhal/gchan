package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
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

type Board struct {
	Board       string `json:"board"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
}

type Thread struct {
	ID       int    `json:"id"`
	Board    string `json:"board"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Usermode string `json:"usermode"`
	Post     string `json:"post"`
	Files    string `json:"files"`
	Created  string `json:"created"`
}

type Threads struct {
	Threads []Thread `json:"threads"`
}

// checkErr function for error handling
func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// open db connection
	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)

	defer db.Close()
	query, err := db.Query("select board from boards")
	for query.Next() {
		var board string
		query.Scan(&board)
		fmt.Printf(board)

	}

}
func boardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["BOARD"]
	// open db connection
	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)

	defer db.Close()
	query, err := db.Query("select title, subtitle, description from boards where board=?", html.EscapeString(board))
	checkErr(err)
	for query.Next() {
		var title, subtitle, description string
		query.Scan(&title, &subtitle, &description)

		fmt.Printf(board)
		fmt.Printf(title)
		fmt.Printf(subtitle)
		fmt.Printf(description)
	}
	stmt, err := db.Query("select * from threads where board=?", html.EscapeString(board))
	checkErr(err)

	b := Threads{Threads: []Thread{}}
	for stmt.Next() {
		p := Thread{}
		stmt.Scan(&p.ID, &p.Board, &p.Title, &p.Name, &p.Email, &p.Usermode, &p.Post, &p.Files, &p.Created)
		b.Threads = append(b.Threads, p)

	}

}

func threadHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//appnumber := vars["appnumber"]
}

func newHandler(w http.ResponseWriter, r *http.Request) {

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
	router.HandleFunc("/{BOARD}/", boardHandler)
	router.HandleFunc("/{BOARD}/thread/{ID}", rootHandler)
	router.HandleFunc("/thread/{ID}", threadHandler)
	router.HandleFunc("/{BOARD}/page/{ID}", boardHandler)
	router.HandleFunc("/img/{ID}", threadHandler)
	router.HandleFunc("/new", newHandler)
	// ListenAndServe on PORT with router
	err = http.ListenAndServe(configuration.Port, router)
	if err != nil {
		log.Fatal(err)
	}

}
