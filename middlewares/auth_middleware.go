package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/golang-jwt-mysql/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		clientToken := ginContext.Request.Header.Get("token")
		if clientToken == "" {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			ginContext.Abort()
			return
		}

		claims, error := helpers.ValidateToken(clientToken)
		if error != "" {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": error})
			ginContext.Abort()
			return
		}

		ginContext.Set("first_name", claims.FirstName)
		ginContext.Set("last_name", claims.LastName)
		ginContext.Set("email", claims.Email)
		ginContext.Set("user_id", claims.UserID)
		ginContext.Set("user_type", claims.UserType)
		ginContext.Next()
	}
}
