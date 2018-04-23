package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Todo Structure
type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Category string `json:"category"`
	IsDone   string `json:"isdone"`
}

/* Globals variables */
var db *gorm.DB
var err error
var todo Todo

// connect to db
func connect() {
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("Failed to connect to database")
	}
}

// InitialMigration for DB
func InitialMigration() {
	connect()
	db.AutoMigrate(&Todo{})
}

// SetHeaders handler
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// GetTodos hanlder
func GetTodos(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	var todos []Todo
	db.Find(&todos)
	SetHeaders(w)

	json.NewEncoder(w).Encode(todos)
}

// NewTodo handler
func NewTodo(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()
	// Check for request Body
	if r.Body == nil {
		http.Error(w, "Please send body", 400)
		return
	}

	// decode r.Body
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}
	db.Create(&Todo{Title: todo.Title, Category: todo.Category, IsDone: todo.IsDone})

	// encode to json
	val, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	SetHeaders(w)

	w.Write(val)
}

func parseBody(w http.ResponseWriter, r *http.Request) {
	// Check for request Body
	if r.Body == nil {
		http.Error(w, "Please send body", 400)
		return
	}
	// Decode todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		panic(err)
	}

	value, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	SetHeaders(w)

	w.Write(value)

	fmt.Println(todo)
}

func main() {
	InitialMigration()

	http.HandleFunc("/", parseBody)
	http.HandleFunc("/todos", GetTodos)
	http.HandleFunc("/create", NewTodo)

	fmt.Println("Server is running at port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
