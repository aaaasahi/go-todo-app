package controllers

import (
	"fmt"
	"go-todo-app/models"
	"go-todo-app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// コントローラ構造体：サービスへの依存を内部に保持
type TodoController struct {
	service *services.TodoService
}

// コンストラクタ関数：サービス構造体を受け取ってコントローラを生成
func NewTodoController(s *services.TodoService) *TodoController {
	return &TodoController{service: s}
}

func (ctrl *TodoController) GetTodosHandler(c echo.Context) error {
	todos, err := ctrl.service.ListTodosService()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch todos")
	}
	return c.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	todo, err := ctrl.service.GetTodoService(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch todo")
	}

	return c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) CreateTodoHandler(c echo.Context) error {
	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	createdTodo, err := ctrl.service.CreateTodoService(reqTodo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.JSON(http.StatusCreated, createdTodo)
}

func (ctrl *TodoController) UpdateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	reqTodo.ID = id
	updatedTodo, err := ctrl.service.UpdateTodoService(reqTodo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update todo")
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func (ctrl *TodoController) DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	if err := ctrl.service.DeleteTodoService(id); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete todo")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Todo with ID %d has been deleted", id),
	})
}
