package controller

import (
	"my-joke-book/database"
	"my-joke-book/database/models"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	c.String(200, "Hello World!!")
}
