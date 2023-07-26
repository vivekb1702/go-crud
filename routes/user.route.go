package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/", controllers.CreateUser)
	r.GET("/:email", controllers.GetUser)
	r.GET("/", controllers.GetUsers)
	r.DELETE("/:email", controllers.DeleteUser)
	r.PATCH("/", controllers.UpdateUser)
}
