package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	databases "github.com/gupta29470/golang-jwt-mysql/database"
	"github.com/gupta29470/golang-jwt-mysql/helpers"
	"github.com/gupta29470/golang-jwt-mysql/models"
	"gorm.io/gorm"
)

var validate *validator.Validate = validator.New()
var db *gorm.DB = databases.DBInstance()

func Signup() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var user models.User

		bindJSONError := ginContext.BindJSON(&user)
		if bindJSONError != nil {
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": bindJSONError.Error()})
			return
		}

		validationError := validate.Struct(user)
		if validationError != nil {
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		password := helpers.HashPasswordHelper(user.Password)
		user.Password = password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UserID = helpers.Generate16CharUUID()
		token, refreshToken := helpers.GenerateAllTokens(user.FirstName, user.LastName, user.Email, user.UserID, user.UserType)
		user.Token = token
		user.RefreshToken = refreshToken

		transaction := db.Create(&user)
		if transaction.Error != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
			return
		}

		ginContext.JSON(http.StatusOK, &user)
	}
}

func Login() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var user models.User

		bindJSONError := ginContext.BindJSON(&user)
		if bindJSONError != nil {
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": bindJSONError.Error()})
			return
		}

		var foundUser models.User

		transaction := db.Where("email=?", user.Email).First(&foundUser)
		if transaction.Error != nil {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": "You are not signed up"})
			return
		}

		isValidPassword := helpers.VerifyPassword(foundUser.Password, user.Password)
		if !isValidPassword {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Incorrect password"})
			return
		}

		token, refreshToken := helpers.GenerateAllTokens(foundUser.FirstName, foundUser.LastName, foundUser.Email, foundUser.UserID, foundUser.UserType)
		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updateTransaction := db.Model(&foundUser).Where("email=?", foundUser.Email).Updates(&models.User{Token: token, RefreshToken: refreshToken, UpdatedAt: updatedAt})
		if updateTransaction.Error != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
			return
		}

		ginContext.JSON(http.StatusOK, foundUser)
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		requestRefreshToken := ginContext.Request.Header.Get("refresh_token")
		var foundUser models.User

		transaction := db.Where("refresh_token=?", requestRefreshToken).First(&foundUser)
		if transaction.Error != nil || foundUser.UserID == "" {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": "Invalid refresh token"})
			return
		}

		token, refreshToken := helpers.GenerateAllTokens(foundUser.FirstName, foundUser.LastName, foundUser.Email, foundUser.UserID, foundUser.UserType)
		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateTransaction := db.Model(&foundUser).Where("email=?", foundUser.Email).Updates(&models.User{Token: token, RefreshToken: refreshToken, UpdatedAt: updatedAt})
		if updateTransaction.Error != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Refresh token updation failed"})
			return
		}

		foundUser.Token = token
		foundUser.RefreshToken = refreshToken
		foundUser.UpdatedAt = updatedAt

		ginContext.JSON(http.StatusOK, foundUser)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		userID := ginContext.Param("user_id")
		var foundUser models.User

		transaction := db.Where("user_id=?", userID).First(&foundUser)
		if transaction.Error != nil {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": "You are not signed up"})
			return
		}

		if foundUser.UserType == "user" {
			ginContext.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "You are not allowed to do this operation"})
			return
		}

		var users []models.User
		db.Find(&users)

		ginContext.JSON(http.StatusOK, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		userID := ginContext.Param("user_id")
		var foundUser models.User

		transaction := db.Where("user_id=?", userID).First(&foundUser)
		if transaction.Error != nil {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if userID != foundUser.UserID && foundUser.UserType == "user" {
			ginContext.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "You are not allowed to do this operation"})
			return
		}

		ginContext.JSON(http.StatusOK, foundUser)
	}
}
