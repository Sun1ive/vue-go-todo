package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

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