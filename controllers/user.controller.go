package controllers

import (
	"go-crud/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var body struct {
		Name  string
		Email string
	}
	c.Bind(&body)

	user, err := services.CreateUser(body.Name, body.Email)

	if err != nil {
		c.Status(400)
	}

	c.JSON(201, gin.H{
		"data":    user,
		"message": "created",
	})
}

func GetUser(c *gin.Context) {

	email := c.Param("email")
	user, err := services.FindUser(email)

	if user == nil {
		c.Status(404)
		return
	}

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data":    user,
		"message": "get user",
	})
}

func GetUsers(c *gin.Context) {

	user, err := services.FindUsers()

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data":    user,
		"message": "get users",
	})
}
