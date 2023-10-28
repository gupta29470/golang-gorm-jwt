package databases

import (
	"fmt"
	"log"

	"github.com/gupta29470/golang-jwt-mysql/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnWithoutDB = "username:password@tcp(127.0.0.1:3306)/golang_jwt_sql?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func connectDB() *gorm.DB {
	var connectionError error
	db, connectionError = gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if connectionError != nil {
		log.Fatal("Error while connecting to database", connectionError)
	}
	fmt.Println("Connected to database")
	return db
}

func DBInstance() *gorm.DB {
	db := connectDB()
	db.AutoMigrate(&models.User{})
	return db
}
