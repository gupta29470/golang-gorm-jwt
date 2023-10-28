package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gupta29470/golang-jwt-mysql/controllers"
)

func AuthRoutes(engine *gin.Engine) {
	engine.POST("/signup", controllers.Signup())
	engine.POST("/login", controllers.Login())
	engine.POST("refresh/token", controllers.RefreshToken())
}
