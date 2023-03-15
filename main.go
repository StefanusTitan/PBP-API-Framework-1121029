package main

import (
	"echoexplore/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users/:id", controllers.GetAUserDetail)
	e.POST("/users", controllers.InsertNewUser)
	e.PUT("/users", controllers.UpdateAUser)
	e.DELETE("/users", controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
