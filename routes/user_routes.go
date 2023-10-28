package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/golang-jwt-mysql/controllers"
	"github.com/gupta29470/golang-jwt-mysql/middlewares"
)

func UserRoutes(engine *gin.Engine) {
	engine.Use(middlewares.AuthMiddleware())
	engine.GET("/users/:user_id", controllers.GetUsers())
	engine.GET("/user/:user_id/", controllers.GetUser())
}
