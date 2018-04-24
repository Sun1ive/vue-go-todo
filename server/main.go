package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Todo Structure
type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Category string `json:"category"`
	IsDone   string `json:"isDone"`
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

// GetTodos handler
func GetTodos(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	var todos []Todo
	db.Find(&todos)
	SetHeaders(w)

	json.NewEncoder(w).Encode(todos)
}

// CreateTodo handler
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	connect()
	defer db.Close()

	var newTodo Todo
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newTodo)
	if err != nil {
		panic(err)
	}

	val := db.Create(&Todo{Title: newTodo.Title, Category: newTodo.Category, IsDone: newTodo.IsDone})

	SetHeaders(w)

	value, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}

	w.Write(value)

}

// DeleteTodo handler
func DeleteTodo(w http.ResponseWriter, r *http.Request)  {
	connect()
	defer db.Close()

	var todo Todo
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)

	if err != nil {
		panic(err)
	}
	
	db.Where("ID = ?", todo.ID).Find(&todo)
	db.Delete(&todo)

	SetHeaders(w)
	
	fmt.Fprintf(w, "Deleted")
}

func UpdateTodo(w http.ResponseWriter,r *http.Request) {
	connect()
	defer db.Close()

	vars:= mux.Vars(r)
	id := vars["id"]
	isDone := vars["isDone"]

	var todo Todo

	db.Where("ID = ?", id).Find(&todo)

	todo.IsDone = isDone

	db.Save(&todo)

	SetHeaders(w)

	fmt.Fprint(w, "Successfully updated")
}

// HandlerRequests router function
func HandlerRequests() {
	var Router = mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/todos", GetTodos).Methods("GET")
	Router.HandleFunc("/create", CreateTodo).Methods("POST")
	Router.HandleFunc("/delete", DeleteTodo).Methods("DELETE")
	Router.HandleFunc("/update/id={id}&completed={isDone}", UpdateTodo).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", cors.AllowAll().Handler(Router)))
}

func main() {
	fmt.Println("Server is running at port 8081")
	InitialMigration()
	HandlerRequests()
}
