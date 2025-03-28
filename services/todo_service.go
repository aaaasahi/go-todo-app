package services

import (
	"go-todo-app/models"
	"go-todo-app/repositories"
)

type TodoServiceIF interface {
	GetTodoService(id int) (models.Todo, error)
	ListTodosService() ([]models.Todo, error)
	CreateTodoService(todo models.Todo) (models.Todo, error)
	UpdateTodoService(todo models.Todo) (models.Todo, error)
	DeleteTodoService(id int) error
}

func (s *TodoService) GetTodoService(id int) (models.Todo, error) {
	todo, err := repositories.GetTodo(s.db, id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) ListTodosService() ([]models.Todo, error) {
	todos, err := repositories.ListTodos(s.db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) CreateTodoService(todo models.Todo) (models.Todo, error) {
	createdTodo, err := repositories.CreateTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil
}

func (s *TodoService) UpdateTodoService(todo models.Todo) (models.Todo, error) {
	updatedTodo, err := repositories.UpdateTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return updatedTodo, nil
}

func (s *TodoService) DeleteTodoService(id int) error {
	err := repositories.DeleteTodo(s.db, id)
	if err != nil {
		return err
	}

	return nil
}
