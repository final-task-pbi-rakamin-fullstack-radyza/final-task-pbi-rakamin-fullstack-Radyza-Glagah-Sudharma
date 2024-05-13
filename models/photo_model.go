package models

import (
	"project-final/database"
	"project-final/app"
)

func CreatePhoto(photo *app.Photo) error {
	result := database.GetDB().Create(photo)
	return result.Error
}

func FindPhotoByID(photoID uint) (app.Photo, error) {
	var photo app.Photo
	result := database.GetDB().First(&photo, photoID)
	return photo, result.Error
}

func UpdatePhoto(photo *app.Photo) error {
	result := database.GetDB().Save(photo)
	return result.Error
}

func DeletePhoto(photoID uint) error {
	result := database.GetDB().Delete(&app.Photo{}, photoID)
	return result.Error
}
