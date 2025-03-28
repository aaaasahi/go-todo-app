package main

import (
	"fmt"
	"go-todo-app/services"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = "127.0.0.1"
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("MYSQL_DATABASE")
)

func main() {
	// データベースに接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	service := services.NewTodoService(db)

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
