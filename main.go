package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go_quiz_6",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	// make a variable for containing data from slice of struct User
	var users []User

	// query orm find all data coming from users table
	result := DB.Find(&users)

	// check the error if required
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error.Error())
	}

	// return the message and data JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = User{}

	// query find the data based on id
	result := DB.First(&user, id)

	// check the error if required
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error.Error())
	}

	// return the message and data JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// make a variable for containing data from struct User
	var user = User{}
	// data binding from user input (json, form, etc format)
	c.Bind(&user)

	// do save query
	result := DB.Save(&user)

	// check the error if required
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error.Error())
	}

	// return the message and data JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = User{}

	// query for deleting data based on id
	result := DB.Delete(&user, id)

	// check the error if required
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error.Error())
	}

	// return the message and data JSON
	return c.JSON(http.StatusOK, "success delete user")
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")

	// make a variable for containing data from struct User
	var user = User{}

	// get data based on id
	DB.First(&user, id)

	// binding the user input
	c.Bind(&user)

	// query to save data to table
	result := DB.Save(&user)

	// check the error if required
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error.Error())
	}

	// return the message and data JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()

	// Routing
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
