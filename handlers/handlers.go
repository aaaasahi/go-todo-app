package handler

import (
	"io"
	"net/http"
)

func GetTodosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Get all TODOs!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Get TODO with ID 1!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Create new TODO!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UpdateTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		io.WriteString(w, "Update TODO with ID 2!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		io.WriteString(w, "Delete TODO with ID 3!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
