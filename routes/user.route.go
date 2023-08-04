package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	group := r.Group("/user")
	group.POST("/", controllers.CreateUser)
	group.GET("/:email", controllers.GetUser)
	group.GET("/", controllers.GetUsers)
	group.DELETE("/:email", controllers.DeleteUser)
	group.PATCH("/", controllers.UpdateUser)
}
