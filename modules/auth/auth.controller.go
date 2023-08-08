package auth

import (
	"go-crud/utils"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	token, err := LoginUserService(body.Email, body.Password)

	utils.CustomResponse(c, &utils.Response{Message: "user logged in", Data: *token, Status: 201}, err)

}
