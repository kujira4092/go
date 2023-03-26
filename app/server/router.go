package server

import (
	"io"
	"html/template"
	"net/http"
	
	"github.com/labstack/echo/v4"

	"app/controller"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func InitRouter(e *echo.Echo) {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/view/*.html")),
	}
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello", controller.Hello)
}