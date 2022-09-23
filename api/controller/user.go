package controller

import (
	"my-joke-book/database"
	"my-joke-book/database/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUsers(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	database.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	c.JSON(200, &users)
}
