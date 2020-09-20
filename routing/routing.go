package routing

import (
	"github.com/TeplyyMaksim/echo-router-test/ctrl"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func StartRouter () {
	e := echo.New()
	e.Validator = &CustomValidator{ validator: validator.New() }

	e.GET("/", ctrl.IndexWithQuery)

	e.POST("/users", ctrl.SaveUser)
	e.GET("/users/:id", ctrl.GetUser)
	e.PUT("/users/:id", ctrl.UpdateUser)
	e.DELETE("/users/:id", ctrl.DeleteUser)


	e.Logger.Fatal(e.Start(":1323"))
}