package services

import (
	"go-crud/initializers"
	"go-crud/models"
)

func CreateUser(name string, email string) (*models.User, error) {

	user := models.User{Name: name, Email: email}

	result := initializers.DB.Create(&user)

	return &user, result.Error
}

func FindUser(email string) (*models.User, error) {
	user := models.User{}
	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &user, result.Error
}

func FindUsers() (*[]models.User, error) {
	user := []models.User{}
	result := initializers.DB.Find(&user)
	return &user, result.Error
}
