package models

import "github.com/golang-jwt/jwt"

type SignedDetails struct {
	FirstName string
	LastName  string
	Email     string
	UserID    string
	UserType  string
	jwt.StandardClaims
}
