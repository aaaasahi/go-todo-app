package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodosHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get all TODOs!")
}

func GetTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Get TODO with ID 1!")
}

func CreateTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Create new TODO!")
}

func UpdateTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Update TODO with ID 2!")
}

func DeleteTodoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Delete TODO with ID 3!")
}
