package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Message string
	Data    interface{}
	Status  int
}

func CustomResponse(c *gin.Context, content *Response, err error) {

	var isError bool

	if err != nil {
		isError = true
		if customErr, ok := err.(CustomError); ok {

			content.Status = customErr.Code
			content.Message = customErr.Message
		} else {
			content.Status = 500
			content.Message = err.Error()
		}
	}

	res := gin.H{
		"version": "v1",
		"data":    content.Data,
		"message": content.Message,
		"err":     isError,
	}

	c.JSON(content.Status, res)
	return
}