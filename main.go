package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/todos/list", getTodosHandler)
	http.HandleFunc("/todos/1", getTodoHandler)
	http.HandleFunc("/todos", createTodoHandler)
	http.HandleFunc("/todos/2", updateTodoHandler)
	http.HandleFunc("/todos/3", deleteTodoHandler)
	log.Println("server start at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Get all TODOs!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Get TODO with ID 1!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Create new TODO!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func updateTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		io.WriteString(w, "Update TODO with ID 2!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func deleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		io.WriteString(w, "Delete TODO with ID 3!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
