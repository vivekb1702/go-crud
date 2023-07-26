package controllers

import (
	"go-crud/services"
	"go-crud/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}
	c.Bind(&body)

	user, err := services.CreateUser(body.Name, body.Email, body.Password)

	utils.CustomResponse(c, &utils.Response{Message: "user created", Data: user, Status: 201}, err)

}

func GetUser(c *gin.Context) {

	email := c.Param("email")
	user, err := services.FindUser(email)

	utils.CustomResponse(c, &utils.Response{Message: "get user", Data: user, Status: 200}, err)

}

func GetUsers(c *gin.Context) {

	user, err := services.FindUsers()

	if err != nil {
		c.Status(400)
		return
	}

	utils.CustomResponse(c, &utils.Response{Message: "get users", Data: user, Status: 200}, err)

}

func UpdateUser(c *gin.Context) {

	var body struct {
		Name  string
		Email string
	}
	c.Bind(&body)

	user, err := services.UpdateUser(body.Name, body.Email)

	utils.CustomResponse(c, &utils.Response{Message: "user updated", Data: user, Status: 201}, err)
}

func DeleteUser(c *gin.Context) {
	email := c.Param("email")
	user, err := services.DeleteUser(email)

	utils.CustomResponse(c, &utils.Response{Message: "user deleted", Data: user, Status: 204}, err)
}
