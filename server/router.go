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

func HandlerRequests() {
	var Router = mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/todos", GetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", cors.Default().Handler(Router)))
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

func main() {
	fmt.Println("Server is running at port 8081")
	InitialMigration()
	HandlerRequests()
}
