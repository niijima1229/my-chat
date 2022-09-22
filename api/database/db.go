package database

import (
	"fmt"
	"my-joke-book/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
