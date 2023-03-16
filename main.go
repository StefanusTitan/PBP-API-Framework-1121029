package main

import (
	"crypto/subtle"
	"echoexplore/controllers"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
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

	g2 := e.Group("jwt")
	g2.Use(echojwt.JWT([]byte("secret")))

	g.GET("/allusers", controllers.GetAllUserDetails)
	e.GET("/users/:id", controllers.GetAUserDetail)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Logic to verify or check something
			// check
			// IF .....
			// For invalid credentials
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			// return next(c)
		}
	})

	e.POST("/users", controllers.InsertNewUser)
	e.PUT("/users", controllers.UpdateAUser)
	e.DELETE("/users", controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
