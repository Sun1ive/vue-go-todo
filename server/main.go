package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// Todo Model struct
type Todo struct {
	gorm.Model
	Todo  string `json:"todo"`
}

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
func GetTodos(w http.ResponseWriter, r *http.Request)  {
	connect()
	defer db.Close()

	var todos []Todo
	db.Find(&todos)

	setHeaders(w)
	fmt.Println("Hello world", todos)
	json.NewEncoder(w).Encode(todos)
}



func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/todos", GetTodos).Methods("GET")
	// myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	// myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	// myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	fmt.Println("Server is running at port 3000")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}


func main() {
	InitialMigration()
	handleRequests()
}