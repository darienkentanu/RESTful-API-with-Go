package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// ----------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, M{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Echo) error {
	// your solution here
}

func DeleteUserController(c echo.Context) error {
	// your solution here
}

func UpdateUserControllers(c echo.Context) error {
	// your solution here
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, M{
		"messages": "success create user",
		"user":     user,
	})
}
