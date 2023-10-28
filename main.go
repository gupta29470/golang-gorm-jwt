package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/golang-jwt-mysql/routes"
	"github.com/joho/godotenv"
)

func main() {
	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal("Something went wrong while loading env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AuthRoutes(engine)
	routes.UserRoutes(engine)

	engine.Run(":" + port)
}
