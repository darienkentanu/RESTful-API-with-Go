package main

import (
	c "github.com/darienkentanu/RESTful-API-with-Go/controller"
	"github.com/labstack/echo"
)

// --------------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", c.GetUsersController)
	e.POST("/users", c.CreateUserController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users/:id", c.UpdateUserControllers)
	e.DELETE("/users/:id", c.DeleteUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
