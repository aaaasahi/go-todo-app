package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
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
