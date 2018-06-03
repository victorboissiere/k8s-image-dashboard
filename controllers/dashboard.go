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

func Pod(c echo.Context) error {
	pod := api.GetPod(c.Param("namespace"), c.Param("podName"))
	if pod.Name == "NotFound" {
		return c.Redirect(302, "/")
	}

	return c.Render(http.StatusOK, "pod.html",  pod)
}

func ArbitraryRollingUpdateAction(c echo.Context) error {
	api.TriggerArbitraryRollingUpdateOnDeployment(c.Param("namespace"), c.Param("podName"))

	return c.NoContent(http.StatusOK)
}
