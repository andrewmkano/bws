package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

// HealthCheck a simple handler to determine the server's health
func HealthCheck(c buffalo.Context) error {
	c.Logger().Infof("Calling health-check")
	return c.Render(http.StatusOK, nil)
}
