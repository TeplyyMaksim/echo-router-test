package routing

import "github.com/labstack/echo"

func startRouter () {
	e := echo.New()

	e.GET("/", index)
}