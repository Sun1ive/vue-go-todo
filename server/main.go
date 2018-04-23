package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Title string `json:"title"`
	Category string `json:"category"`
	IsDone string `json:"isdone"`
}

func SetHeaders(w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json")
}

func parseBody(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)

	var todo Todo
	err := decoder.Decode(&todo)

	if err != nil {
		panic(err)
	}

	fmt.Println(todo.Title)
	fmt.Println(todo.Category)
	fmt.Println(todo.IsDone)

	output, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	SetHeaders(w)

	w.Write(output)
}

func main()  {
	http.HandleFunc("/", parseBody)
	http.ListenAndServe(":8081", nil)
}