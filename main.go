package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func setup() {

}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")

}
func boardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["board"]
	db, err := sql.Open("mysql", "chan:test@/chan?charset=utf8")
	check(err)
	var uri string
	err = db.QueryRow("select uri from boards where uri = ?", html.EscapeString(board)).Scan(&uri)
	if err == sql.ErrNoRows {
		fmt.Fprintf(w, "Error no board")
	} else {
		fmt.Println(uri)
	}

}
func threadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["board"]
	threadID := vars["threadId"]
	fmt.Fprintf(w, board, threadID)

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["board"]
	threadID := vars["threadId"]
	fmt.Fprintf(w, board, threadID)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	post := r.Form["post"]
	//	image := r.Form["image"]
	fmt.Println(post)

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/new", newPost)
	r.HandleFunc("/{board}", boardHandler)
	r.HandleFunc("/{board}/{threadId}", threadHandler)
	r.HandleFunc("/{board}/{threadId}/{output}", apiHandler)
	r.HandleFunc("/{board}/{threadId}/{output}", apiHandler)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatal(err)
	}

}
