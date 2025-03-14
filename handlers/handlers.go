package handler

import (
	"fmt"
	"go-todo-app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	todos := []models.Todo{models.Todo1, models.Todo2}
	return c.JSON(http.StatusOK, todos)
}

func GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	// 暫定でログを出力
	log.Println(id)

	return c.JSON(http.StatusOK, models.Todo1)
}

func CreateTodoHandler(c echo.Context) error {
	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to parse request body")
	}

	todo := reqTodo

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}

	var reqTodo models.Todo
	if err := c.Bind(&reqTodo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to parse request body")
	}

	todo := reqTodo
	todo.ID = id

	return c.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Delete TODO with ID %d!", id))
}
