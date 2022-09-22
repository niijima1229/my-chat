package routes

import (
	"my-joke-book/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
