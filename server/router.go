package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

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

func InitialMigration() {
	connect()
	defer db.Close()

	db.AutoMigrate(&Todo{})
}



func GetTodos(w http.ResponseWriter, r *http.Request)  {
	connect()
	defer db.Close()

	var todos []Todo
	db.Find(&todos)

	setHeaders(w)
	json.NewEncoder(w).Encode(todos)
}