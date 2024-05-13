package models

import (
	"project-final/database"
	"project-final/app"
)

func CreateUser(user *app.User) error {
	result := database.GetDB().Create(user)
	return result.Error
}

func FindUserByID(userID uint) (app.User, error) {
	var user app.User
	result := database.GetDB().First(&user, userID)
	return user, result.Error
}

func UpdateUser(user *app.User) error {
	result := database.GetDB().Save(user)
	return result.Error
}

func DeleteUser(userID uint) error {
	result := database.GetDB().Delete(&app.User{}, userID)
	return result.Error
}
