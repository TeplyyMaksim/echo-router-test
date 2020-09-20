package ctrl

import (
	"github.com/TeplyyMaksim/echo-router-test/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	User struct {
		Id 		int 	`json:"id" validate:"required"`
		Name 	string	`json:"name"`
	}
)

var database = map[int]*User{
	10: { Id: 10, Name: "Maksym" },
	20: { Id: 20, Name: "MaxMaksym" },
}

func SaveUser (c echo.Context) error {
	var newUser User

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	database[newUser.Id] = &newUser
	return c.JSON(http.StatusOK, newUser)
}

func GetUser (c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:	http.StatusBadRequest,
		})
	}

	user := database[id]

	if user == nil {
		return c.JSON(http.StatusNotFound, utils.Error{
			Message: "User not found",
			Code: http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser (c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:	http.StatusBadRequest,
		})
	}

	user := database[id]

	if user == nil {
		return c.JSON(http.StatusNotFound, utils.Error{
			Message: "User not found",
			Code: http.StatusNotFound,
		})
	}

	var updatedUser User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{
			Message: err.Error(),
			Code:	http.StatusBadRequest,
		})
	}

	user := database[id]

	if user == nil {
		return c.JSON(http.StatusNotFound, utils.Error{
			Message: "User not found",
			Code: http.StatusNotFound,
		})
	}

	database[id] = nil

	return c.JSON(http.StatusOK, struct { Message string `json:"message"` }{ Message: "Successfully deleted" })
}