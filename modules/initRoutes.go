package modules

import (
	"go-crud/modules/auth"
	"go-crud/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	version1 := engine.Group("/api/v1")
	user.UserRoutes(version1)
	auth.AuthRoutes(version1)
}
