package controllers

import (
	"echo-golang-quiz6/configs"
	"echo-golang-quiz6/models"
	"echo-golang-quiz6/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	// make a variable for containing data from slice of struct User
	var users []models.User

	// call repositories to get all user data
	err := repositories.GetUsers(&users)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mendapatkan data",
		Status:  true,
		Data:    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = models.User{}

	// call repositories to get user data based on id
	err := repositories.GetUser(&user, id)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mendapatkan data",
		Status:  true,
		Data:    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// make a variable for containing data from struct User
	var userRequest = models.User{}
	// data binding from user input (json, form, etc format)
	c.Bind(&userRequest)

	// call repositories to create data
	err := repositories.CreateUser(&userRequest)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambahkan data",
		Status:  true,
		Data:    userRequest,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = models.User{}

	// get data based on id
	configs.DB.First(&user, id)

	// binding the user input
	c.Bind(&user)

	// call repositories to update data
	err := repositories.UpdateUser(&user, id)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil merubah data",
		Status:  true,
		Data:    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = models.User{}

	// call repositories to delete data
	err := repositories.DeleteUser(&user, id)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menghapus data",
		Status:  true,
		Data:    nil,
	})
}
