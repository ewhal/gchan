package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func setup() {
	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")

}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")

}
func boardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["board"]
	fmt.Fprintf(w, board)

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
	/*	r.HandleFunc("/{board}", boardHandler)
		r.HandleFunc("/{board}/{threadId}", threadHandler)
		r.HandleFunc("/{board}/{threadId}/{output}", apiHandler)
		r.HandleFunc("/{board}/{threadId}/{output}", apiHandler) */
	//	http.Handle("/", r)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatal(err)
	}

}
