package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type IndexData struct {
	UserName string
}

func renderIndex(c echo.Context, code int, name string, data IndexData) error {
	return c.Render(code, name, data)
}

func index(c echo.Context) error {
	userName := c.QueryParam("name")
	err := renderIndex(c, 200, "index", IndexData{UserName: userName})
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//go:embed templates/*
var HTMLTemplates embed.FS

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseFS(HTMLTemplates, "templates/*.html.tmpl")),
	}

	e.Renderer = t

	e.GET("/", index)
	e.Logger.Debug(e.Start(":1323"))
}
