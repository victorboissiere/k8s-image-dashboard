package dashboard

import (
	"net/http"
	"github.com/labstack/echo"
	"../api"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", api.GetNamespacePods())
}

func Nodes(c echo.Context) error {
	return c.Render(http.StatusOK, "nodes.html", api.GetNodes())
}
