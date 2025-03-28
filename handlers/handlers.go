package handler

import (
	"fmt"
	"go-todo-app/models"
	"go-todo-app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	todos, err := services.ListTodosService()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch todos")
	}
	return c.JSON(http.StatusOK, todos)
}

func GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	todo, err := services.GetTodoService(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch todo")
	}

	return c.JSON(http.StatusOK, todo)
}

func CreateTodoHandler(c echo.Context) error {
	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	createdTodo, err := services.CreateTodoService(reqTodo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.JSON(http.StatusCreated, createdTodo)
}

func UpdateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	reqTodo.ID = id
	updatedTodo, err := services.UpdateTodoService(reqTodo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update todo")
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	if err := services.DeleteTodoService(id); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete todo")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Todo with ID %d has been deleted", id),
	})
}
