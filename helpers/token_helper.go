package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gupta29470/golang-jwt-mysql/models"
)

const (
	secretKey = "HwfoReq9lKSvLm+ONuAQ21JFkMeFx3JxqngToE48D4s="
)

func GenerateAllTokens(firstName string, lastName string, email string, userID string, userType string) (string, string) {
	claims := &models.SignedDetails{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserID:    userID,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(124)).Unix(),
		},
	}

	refreshClaims := &models.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, tokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, refreshTokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if tokenError != nil || refreshTokenError != nil {
		log.Panic(tokenError, refreshTokenError)
	}

	return token, refreshToken
}

func ValidateToken(clientToken string) (*models.SignedDetails, string) {
	signedDetails := &models.SignedDetails{}
	token, error := jwt.ParseWithClaims(clientToken,
		signedDetails,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if error != nil {
		return nil, error.Error()
	}

	if !token.Valid {
		return nil, "token is invalid"
	}

	return signedDetails, ""
}
