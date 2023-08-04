package modules

import (
	"go-crud/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	user.UserRoutes(engine)
}
