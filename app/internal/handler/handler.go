package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
}

func NewHandler(e *echo.Echo) {
	handler := &handler{}
	e.GET("/test", handler.Test)
}

type Handler interface {
	Test(c echo.Context) error
}

func (h *handler) Test(c echo.Context) error {

	response := "success"

	return c.JSON(http.StatusOK, response)
}
