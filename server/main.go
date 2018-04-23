package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// Todo Model struct
type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Category string `json:"category"`
	IsDone   string `json:"isDone"`
}

// TODOS LIST
var todos []Todo

func setHeaders(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func connect() {
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("Failed to connect to database")
	}
}

// InitialMigration db
func InitialMigration() {
	connect()
	defer db.Close()

	db.AutoMigrate(&Todo{})
}

// GetTodos Endpoint
func GetTodos(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	db.Find(&todos)

	setHeaders(w)
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo handler
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// connect()
	// defer db.Close()

	// vars := mux.Vars(r)
	// title := vars["title"]
	// category := vars["category"]
	// isDone := vars["isDone"]

	// db.Create(&Todo{Title: title, Category: category, IsDone: isDone})
	// setHeaders(w)
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Println(title)
	// fmt.Fprintf(w, "TODO CREATED")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/todos", GetTodos).Methods("GET")
	myRouter.HandleFunc("/create/{title}", CreateTodo).Methods("POST")
	// myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	// myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	fmt.Println("Server is running at port 3000")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	InitialMigration()
	handleRequests()
}
