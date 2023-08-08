package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	group := r.Group("/user")
	group.POST("/", CreateUser)
	group.GET("/:email", GetUser)
	group.GET("/", GetUsers)
	group.DELETE("/:email", DeleteUser)
	group.PATCH("/", UpdateUser)
}
