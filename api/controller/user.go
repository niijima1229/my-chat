package controller

import (
	"my-joke-book/database"
	"my-joke-book/database/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := database.ConnectDB()
	users := []models.User{}
	db.Find(&users)
	c.JSON(200, &users)
}

func CreateUsers(c *gin.Context) {
	var user models.User
	db := database.ConnectDB()
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	db := database.ConnectDB()
	db.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	users := []models.User{}
	db := database.ConnectDB()
	db.Find(&users)
	c.JSON(200, &users)
}
