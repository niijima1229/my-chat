package routes

import (
	"my-joke-book/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/api/user", controller.GetUsers)
	router.POST("/api/user", controller.CreateUsers)
	router.DELETE("/api/user/:id", controller.DeleteUser)
	router.PUT("/api/user/:id", controller.UpdateUser)
}
