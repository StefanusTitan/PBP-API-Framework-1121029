package main

import (
	"crypto/subtle"
	"echoexplore/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	g := e.Group("admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("Titan")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("singasari12")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	e.GET("/users/:id", controllers.GetAUserDetail)
	e.POST("/users", controllers.InsertNewUser)
	e.PUT("/users", controllers.UpdateAUser)
	e.DELETE("/users", controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
