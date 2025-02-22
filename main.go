package main

import (
	handler "go-todo-app/handlers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/todos/list", handler.GetTodosHandler)
	e.GET("/todos/1", handler.GetTodoHandler)
	e.POST("/todos", handler.CreateTodoHandler)
	e.PUT("/todos/2", handler.UpdateTodoHandler)
	e.DELETE("/todos/3", handler.DeleteTodoHandler)

	log.Println("server start at port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
