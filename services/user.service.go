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

type findUserResponse struct {
	ID    uint
	Email string
	Name  string
}

func FindUser(email string) (*findUserResponse, error) {

	var response findUserResponse

	user := models.User{}
	result := initializers.DB.Model(&user).Where("email = ?", email).Select("ID", "Name", "Email").First(&response)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return &response, result.Error
}

func FindUsers() (*[]findUserResponse, error) {
	var response []findUserResponse
	user := models.User{}
	result := initializers.DB.Model(&user).Select("ID", "Name", "Email").Find(&response)
	return &response, result.Error
}

func DeleteUser(email string) (*models.User, error) {

	user := models.User{}

	findUser := initializers.DB.Where("email = ?", email).First(&user)

	if findUser.RowsAffected == 0 {
		return nil, findUser.Error
	}

	result := initializers.DB.Delete(&user, user.ID)

	return &user, result.Error

}
