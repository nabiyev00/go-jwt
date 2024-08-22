package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nabiyev00/go-jwt/controller"
	"github.com/nabiyev00/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabse()
}
func main() {
	r := gin.Default()

	r.POST("/signup", controller.SignUp)
	r.POST("/signup", controller.Login)
	r.Run()
}
