package main

import (
	handler "go-todo-app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos/list", handler.GetTodosHandler)
	http.HandleFunc("/todos/1", handler.GetTodoHandler)
	http.HandleFunc("/todos", handler.CreateTodoHandler)
	http.HandleFunc("/todos/2", handler.UpdateTodoHandler)
	http.HandleFunc("/todos/3", handler.DeleteTodoHandler)
	log.Println("server start at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
