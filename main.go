package main

import (
	"go-crud/initializers"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}
func main() {

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()
}
