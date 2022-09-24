package database

import (
	"fmt"
	"my-joke-book/database/models"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	}
	db.AutoMigrate(&models.User{})
	hash, err := HashPassword("password")
	if err != nil {
		fmt.Println("Hash error(Init): ", err)
	}
	var users = []models.User{
		{Name: "jinzhu1", Email: "test@com", Password: hash},
		{Name: "jinzhu2", Email: "test2@com", Password: hash},
		{Name: "jinzhu3", Email: "test3@com", Password: hash}}
	db.Create(&users)
	DB = db
}
