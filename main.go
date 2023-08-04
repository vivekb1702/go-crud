package main

import (
	"go-crud/initializers"
	"go-crud/modules"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}
func main() {

	engine := gin.Default()

	engine.SetTrustedProxies(nil)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	modules.InitRoutes(engine)
	engine.Run()
}
