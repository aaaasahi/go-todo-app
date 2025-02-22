package main

import (
	handler "go-todo-app/handlers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/todos", handler.GetTodosHandler)
	e.GET("/todos/:id", handler.GetTodoHandler)
	e.POST("/todos", handler.CreateTodoHandler)
	e.PUT("/todos/:id", handler.UpdateTodoHandler)
	e.DELETE("/todos/:id", handler.DeleteTodoHandler)

	log.Println("server start at port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
