package main

import (
	"fmt"
	"go-todo-app/controllers"
	"go-todo-app/router"
	"go-todo-app/services"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
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
	// サービス構造体の初期化（DBを注入）
	service := services.NewTodoService(db)
	// コントローラ構造体の初期化（サービスを注入）
	controller := controllers.NewTodoController(service)

	// ルータ層にルーティング構築を委譲
	// ルータ層でEchoとルート設定を生成
	e := router.NewRouter(controller)

	log.Println("server start at port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
