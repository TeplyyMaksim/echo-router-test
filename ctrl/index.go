package ctrl

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func IndexWithQuery(c echo.Context) error {
	firstName := c.QueryParam("firstName")
	lastName := c.QueryParam("lastName")

	return c.JSON(http.StatusOK, struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}{
		FirstName: firstName,
		LastName: lastName,
	})
}