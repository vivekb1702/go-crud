package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.RouterGroup) {
	group := r.Group("/auth")
	group.POST("/login", LoginUser)
}
