package services

import (
	"go-todo-app/models"
	"go-todo-app/repositories"
)

func GetTodoService(id int) (models.Todo, error) {
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

func ListTodosService() ([]models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	todos, err := repositories.ListTodos(db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodoService(todo models.Todo) (models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return models.Todo{}, err
	}
	defer db.Close()

	createdTodo, err := repositories.CreateTodo(db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func UpdateTodoService(todo models.Todo) (models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return models.Todo{}, err
	}
	defer db.Close()

	updatedTodo, err := repositories.UpdateTodo(db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return updatedTodo, nil
}

func DeleteTodoService(id int) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = repositories.DeleteTodo(db, id)
	if err != nil {
		return err
	}

	return nil
}
