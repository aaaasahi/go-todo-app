package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get all TODOs!")
}

func GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Get TODO with ID %d!", id))
}

func CreateTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Create new TODO!")
}

func UpdateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Update TODO with ID %d!", id))
}

func DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID format")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Delete TODO with ID %d!", id))
}
