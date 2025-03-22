package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	dbUser     = "docker"
	dbPassword = "password"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
	dbName     = "testdb"
)

func main() {
	// e := echo.New()

	// e.GET("/todos", handler.GetTodosHandler)
	// e.GET("/todos/:id", handler.GetTodoHandler)
	// e.POST("/todos", handler.CreateTodoHandler)
	// e.PUT("/todos/:id", handler.UpdateTodoHandler)
	// e.DELETE("/todos/:id", handler.DeleteTodoHandler)

	// log.Println("server start at port 8080")
	// if err := e.Start(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

	// データベースに接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to the database")
}
