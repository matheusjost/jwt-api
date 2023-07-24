package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusjost/jwt-api/controllers"
	"github.com/matheusjost/jwt-api/initializers"
	"github.com/matheusjost/jwt-api/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateJWTToken)
	r.Run()
}
