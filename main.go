package main

import (
	"database/sql"
	"encoding/json"
	"html"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var DATABASE string
var configuration Configuration
var templates = template.Must(template.ParseFiles("templates/index.html", "templates/board.html", "templates/thread.html"))

type Configuration struct {
	// PORT for golang to listen on
	Port string
	// USERNAME database username
	Username string
	// PASS database password
	Password string
	// NAME database name
	Name           string
	UploadLocation string
}

type Board struct {
	Board       string `json:"board"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
}

type Thread struct {
	ID        int    `json:"id"`
	Board     string `json:"board"`
	Threadnum int    `json:"threadnum"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Usermode  string `json:"usermode"`
	Post      string `json:"post"`
	Files     string `json:"files"`
	Created   string `json:"created"`
}
type Post struct {
	ID       int    `json:"id"`
	Board    string `json:"board"`
	Postnum  int    `json:"postnum"`
	Thread   int    `json:"thread"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Usermode string `json:"usermode"`
	Post     string `json:"post"`
	Files    string `json:"files"`
	Created  string `json:"created"`
}

type Threads struct {
	Boards  []Board  `json:"boards"`
	Board   Board    `json:"board"`
	Threads []Thread `json:"threads"`
}
type Posts struct {
	Boards  []Board  `json:"boards"`
	Board   Board    `json:"board"`
	Threads []Thread `json:"threads"`
	Posts   []Post   `json:"posts"`
}
type Boards struct {
	Boards []Board `json:"boards"`
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
	b := Boards{Boards: []Board{}}
	query, err := db.Query("select board from boards")
	for query.Next() {
		p := Board{}
		query.Scan(&p.Board)
		b.Boards = append(b.Boards, p)

	}
	err = templates.ExecuteTemplate(w, "index.html", &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func boardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["BOARD"]
	// open db connection
	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)

	defer db.Close()
	b := Posts{Board: Board{}, Boards: []Board{}, Threads: []Thread{}, Posts: []Post{}}
	query, err := db.Query("select board, title, subtitle, description from boards where board=?", html.EscapeString(board))
	checkErr(err)
	for query.Next() {
		query.Scan(&b.Board.Board, &b.Board.Title, &b.Board.Subtitle, &b.Board.Description)

	}
	query, err = db.Query("select board from boards")
	checkErr(err)
	for query.Next() {
		p := Board{}
		query.Scan(&p.Board, &p.Title, &p.Subtitle, &p.Description)
		b.Boards = append(b.Boards, p)

	}
	stmt, err := db.Query("select id, board, threadnum, title, name, email, usermode, post, files, created from threads where board=? LIMIT 15", html.EscapeString(board))
	checkErr(err)

	for stmt.Next() {
		p := Thread{}
		stmt.Scan(&p.ID, &p.Board, &p.Threadnum, &p.Title, &p.Name, &p.Email, &p.Usermode, &p.Post, &p.Files, &p.Created)
		b.Threads = append(b.Threads, p)

	}

	err = templates.ExecuteTemplate(w, "board.html", &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func threadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["BOARD"]
	thread := vars["ID"]

	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)

	defer db.Close()
	b := Posts{Boards: []Board{}, Threads: []Thread{}, Posts: []Post{}}
	query, err := db.Query("select board, title, subtitle, description from boards where board=?", html.EscapeString(board))
	checkErr(err)
	for query.Next() {
		p := Board{}
		query.Scan(&p.Board, &p.Title, &p.Subtitle, &p.Description)
		b.Boards = append(b.Boards, p)

	}
	stmt, err := db.Query("select id, board, threadnum, title, name, email, usermode, post, files, created from threads where board=? LIMIT 1", html.EscapeString(board))
	checkErr(err)

	for stmt.Next() {
		p := Thread{}
		stmt.Scan(&p.ID, &p.Board, &p.Threadnum, &p.Title, &p.Name, &p.Email, &p.Usermode, &p.Post, &p.Files, &p.Created)
		b.Threads = append(b.Threads, p)

	}
	sm, err := db.Query("select id, board, postnum, thread, title, name, email, usermode, post, files, created from posts where board=? and thread=?", html.EscapeString(board), html.EscapeString(thread))
	checkErr(err)

	for sm.Next() {
		p := Post{}
		sm.Scan(&p.ID, &p.Board, &p.Postnum, &p.Thread, &p.Title, &p.Name, &p.Email, &p.Usermode, &p.Post, &p.Files, &p.Created)
		b.Posts = append(b.Posts, p)

	}
	err = templates.ExecuteTemplate(w, "thread.html", &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board := vars["BOARD"]

	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)

	defer db.Close()

	name := r.FormValue("name")
	email := r.FormValue("email")
	post := r.FormValue("post")
	usermode := r.FormValue("usermode")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("files")
	checkErr(err)
	defer file.Close()
	crutime := time.Now().Unix()
	extName := filepath.Ext(handler.Filename)
	// create new filename with random name and extension
	filename := string(crutime) + extName
	// create a new file ready for user to upload to
	f, err := os.OpenFile(configuration.UploadLocation+filename, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(err)
	defer f.Close()
	io.Copy(f, file)

	query, err := db.Prepare("insert into threads(board, thread, name, email, post, thread, usermode, file) values(?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)

	_, err = query.Exec(html.EscapeString(board), time.Now().Unix(), html.EscapeString(name), html.EscapeString(email), html.EscapeString(post), html.EscapeString(usermode), html.EscapeString(filename))
	checkErr(err)

}
func logoutHandler(w http.ResponseWriter, r *http.Request) {

}
func registerHandler(w http.ResponseWriter, r *http.Request) {

}
func loginHandler(w http.ResponseWriter, r *http.Request) {

}
func modHandler(w http.ResponseWriter, r *http.Request) {

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
	router.HandleFunc("/{BOARD}", boardHandler)
	router.HandleFunc("/{BOARD}/", boardHandler)
	router.HandleFunc("/{BOARD}/thread/{ID}", threadHandler)
	router.HandleFunc("/{BOARD}/page/{ID}", boardHandler)
	router.HandleFunc("/img/{ID}", threadHandler)
	router.HandleFunc("/{BOARD}/new", newHandler).Methods("POST")
	router.HandleFunc("/mod", modHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/register", registerHandler)
	router.HandleFunc("/logout", logoutHandler)

	// ListenAndServe on PORT with router
	err = http.ListenAndServe(configuration.Port, router)
	if err != nil {
		log.Fatal(err)
	}

}
