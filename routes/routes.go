package routes

import (
	"internship-stikom/controllers"
	"internship-stikom/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware())
	// protected.GET("/profile",controllers.GetProfile)
	// protected.POST("/project",controllers.CreateProject)
	return r
}
