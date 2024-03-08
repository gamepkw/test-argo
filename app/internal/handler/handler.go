package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
}

func NewHandler(e *echo.Echo) {
	handler := &handler{}
	e.GET("/test", handler.Test)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}

type Handler interface {
	Test(c echo.Context) error
}

func (h *handler) Test(c echo.Context) error {

	response := "success"

	fmt.Println(response)

	return c.JSON(http.StatusOK, response)
}
