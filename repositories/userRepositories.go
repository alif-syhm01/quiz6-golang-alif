package repositories

import (
	"echo-golang-quiz6/configs"
	"echo-golang-quiz6/models"
)

func GetUsers(usersList *[]models.User) error {
	// query find the user then append to user struct
	result := configs.DB.Find(&usersList)

	// check the error
	if result.Error != nil {
		return result.Error
	}

	// return nil if ok
	return nil
}

func GetUser(user *models.User, id string) error {
	// query find the user based on id then append to user struct
	result := configs.DB.First(&user, id)

	// check the error
	if result.Error != nil {
		return result.Error
	}

	// return nil if ok
	return nil
}

func CreateUser(userRequest *models.User) error {
	// query create the user then append to user struct
	result := configs.DB.Create(&userRequest)

	// check the error
	if result.Error != nil {
		return result.Error
	}

	// return nil if ok
	return nil
}

func UpdateUser(userUpdate *models.User, id string) error {
	// query to save data to table
	result := configs.DB.Save(&userUpdate)

	// check the error
	if result.Error != nil {
		return result.Error
	}

	// return nil if ok
	return nil
}

func DeleteUser(userDelete *models.User, id string) error {
	// query to delete data
	result := configs.DB.Delete(&userDelete, id)

	// check the error
	if result.Error != nil {
		return result.Error
	}

	// return nil if ok
	return nil
}
