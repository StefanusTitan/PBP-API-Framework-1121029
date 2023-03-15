package controllers

import (
	"echoexplore/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertNewUser(c echo.Context) error {
	db := gormConn()
	var user model.Users

	user.UserName = c.FormValue("username")
	user.UserEmail = c.FormValue("email")
	user.UserCountry = c.FormValue("country")
	user.UserPassword = c.FormValue("password")

	insertUser := db.Create(&user)
	var response model.Response

	if insertUser.Error == nil {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"
		return c.JSON(http.StatusInternalServerError, response)
	}
}

func GetAUserDetail(c echo.Context) error {
	db := gormConn()
	id := c.Param("id")

	var user model.Users
	var response model.Response

	getUser := db.Where("userId=?", id).First(&user)

	if getUser.Error == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"
		return c.JSON(http.StatusInternalServerError, response)
	}
}

func UpdateAUser(c echo.Context) error {
	db := gormConn()
	var beforeUser model.Users
	var user model.Users

	user.UserId, _ = strconv.Atoi(c.FormValue("id"))
	user.UserName = c.FormValue("username")
	user.UserEmail = c.FormValue("email")
	user.UserCountry = c.FormValue("country")
	user.UserPassword = c.FormValue("password")

	getUser := db.Where("userId=?", user.UserId).First(&beforeUser)
	if getUser.Error != nil {
		log.Fatal(getUser.Error)
	}
	if user.UserName == "" {
		user.UserName = beforeUser.UserName
	} else if user.UserEmail == "" {
		user.UserEmail = beforeUser.UserEmail
	} else if user.UserCountry == "" {
		user.UserCountry = beforeUser.UserCountry
	} else if user.UserPassword == "" {
		user.UserPassword = beforeUser.UserPassword
	} else if user.UserType == 0 {
		user.UserType = beforeUser.UserType
	}

	updateUser := db.Save(&user)
	var response model.Response
	if updateUser.Error == nil {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"
		return c.JSON(http.StatusInternalServerError, response)
	}
}

func DeleteUser(c echo.Context) error {
	db := gormConn()
	var user model.Users
	user.UserId, _ = strconv.Atoi(c.QueryParam("id"))

	deleteUser := db.Delete(&user)
	var response model.Response
	if deleteUser.Error == nil {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"
		return c.JSON(http.StatusInternalServerError, response)
	}
}

func GetAllUserDetails(c echo.Context) error {
	db := gormConn()
	var users []model.Users

	getUsers := db.Find(&users)

	var response model.UsersResponse
	if getUsers.Error == nil {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = users
		return c.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"
		return c.JSON(http.StatusInternalServerError, response)
	}
}
