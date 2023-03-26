package controller

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello.html", "unti")
}