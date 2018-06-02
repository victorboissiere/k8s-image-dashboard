package main

import (
	"github.com/labstack/echo"
	"io"
	"html/template"
	"./controllers"
	"os"
	"net/http"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", dashboard.Index)
	e.GET("/nodes", dashboard.Nodes)
	e.GET("/pod/:namespace/:podName", dashboard.Pod)
	e.POST("/pod/actions/rolling-update/:namespace/:podName", dashboard.Pod)
	e.GET("/health", func (c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
