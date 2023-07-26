package services

import (
	"go-crud/initializers"
	"go-crud/models"
	"go-crud/utils"
)

func CreateUser(name string, email string) (*models.User, error) {

	user := models.User{Name: name, Email: email}

	findUser := initializers.DB.Where("email = ?", email).First(&user)

	if findUser.RowsAffected != 0 {
		return nil, utils.CustomError{Code: 409, Message: "User already exists"}
	}

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
		return nil, utils.CustomError{Code: 404, Message: "User not found"}
	}

	return &response, result.Error
}

func FindUsers() (*[]findUserResponse, error) {
	var response []findUserResponse
	user := models.User{}
	result := initializers.DB.Model(&user).Select("ID", "Name", "Email").Find(&response)
	return &response, result.Error
}

func UpdateUser(name string, email string) (*models.User, error) {

	user := models.User{}

	findUser := initializers.DB.Where("email = ?", email).First(&user)

	if findUser.RowsAffected == 0 {
		return nil, utils.CustomError{Code: 404, Message: "User not found"}
	}

	user.Email = email
	user.Name = name

	result := initializers.DB.Save(&user)

	return &user, result.Error
}

func DeleteUser(email string) (*models.User, error) {

	user := models.User{}

	findUser := initializers.DB.Where("email = ?", email).First(&user)

	if findUser.RowsAffected == 0 {
		return nil, utils.CustomError{Code: 404, Message: "User not found"}
	}

	result := initializers.DB.Delete(&user, user.ID)

	return &user, result.Error

}
