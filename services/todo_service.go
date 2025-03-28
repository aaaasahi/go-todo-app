package services

import (
	"go-todo-app/models"
	"go-todo-app/repositories"
)

// 指定 ID のTODOを返却
func GetTodoService(id int) (models.Todo, error) {
	// TODO : sql.DB 型を手に入れて、変数 db に代入する
	db, err := connectDB()
	if err != nil {
		return models.Todo{}, err
	}
	defer db.Close()

	todo, err := repositories.GetTodo(db, id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}
